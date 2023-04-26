package repository

import (
	"errors"
	"io/ioutil"
	"math"
	"mime/multipart"
	"os"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
	"github.com/gin-gonic/gin"
	"github.com/slyxh2/golang-blog/interfaces"
	"github.com/slyxh2/golang-blog/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type postRepository struct {
	database   *mongo.Database
	collection string
	awsSession *session.Session
}

func NewPostRepository(db *mongo.Database) (*postRepository, error) {
	cred := credentials.NewStaticCredentials(os.Getenv("BUCKET_KEY_ID"), os.Getenv("BUCKET_SECRET_KEY"), "")
	// The session the S3 Uploader will use
	sess, err := session.NewSession(&aws.Config{
		Region:      aws.String("us-east-1"),
		Credentials: cred,
	})
	if err != nil {
		return nil, err
	}
	return &postRepository{
		database:   db,
		collection: interfaces.CollectionPost,
		awsSession: sess,
	}, nil
}

func (pr *postRepository) Upload(c *gin.Context, file multipart.File, post *models.Post, categoryId string) (*s3manager.UploadOutput, error) {
	if file == nil {
		return nil, errors.New("file is nil")
	}
	if categoryId == "" {
		return nil, errors.New("lack category id")
	}
	// Create an uploader with the session and default options
	uploader := s3manager.NewUploader(pr.awsSession)
	bucket := os.Getenv("BUCKET_NAME")
	key := post.Id.Hex() + ".md"

	// Upload the file to S3.
	result, err := uploader.Upload(&s3manager.UploadInput{
		Bucket: aws.String(bucket),
		Key:    aws.String(key),
		Body:   file,
	})
	if err != nil {
		return nil, err
	}

	var category models.Category
	categoryCollection := pr.database.Collection(interfaces.CollectionCategory)
	objID, err := primitive.ObjectIDFromHex(categoryId)
	if err != nil {
		return nil, err
	}
	err = categoryCollection.FindOneAndUpdate(c, bson.M{"_id": objID}, bson.M{"$push": bson.M{"posts": post}}).Decode(&category)
	if err != nil {
		return nil, err
	}
	post.Category = category

	collection := pr.database.Collection(pr.collection)
	_, err = collection.InsertOne(c, post)
	return result, err
}

func (pr *postRepository) DownLoad(id string) (string, error) {
	awsClient := s3.New(pr.awsSession)
	bucket := os.Getenv("BUCKET_NAME")
	key := id + ".md"

	resp, err := awsClient.GetObject(&s3.GetObjectInput{
		Bucket: aws.String(bucket),
		Key:    aws.String(key),
	})
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	// Read the contents of the file into a byte slice
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	return string(body), nil
}

func (pr *postRepository) Delete(c *gin.Context, id string) error {
	collection := pr.database.Collection(pr.collection)
	categoryCollection := pr.database.Collection(interfaces.CollectionCategory)
	var post models.Post
	postID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return err
	}
	err = collection.FindOneAndDelete(c, bson.M{"_id": postID}).Decode(&post)
	if err != nil {
		return err
	}
	categoryId := post.Category.Id

	filter := bson.M{"_id": categoryId}

	// define the update to remove the post with the given postId
	update := bson.M{"$pull": bson.M{"posts": bson.M{"_id": postID}}}

	// execute the update operation with the filter and update
	_, err = categoryCollection.UpdateOne(c, filter, update)
	if err != nil {
		return err
	}
	awsClient := s3.New(pr.awsSession)
	bucket := os.Getenv("BUCKET_NAME")
	key := id + ".md"
	_, err = awsClient.DeleteObject(&s3.DeleteObjectInput{
		Bucket: aws.String(bucket),
		Key:    aws.String(key),
	})
	if err != nil {
		return err
	}
	return nil
}

func (pr *postRepository) Edit(c *gin.Context, id string, header string, file multipart.File) error {
	if header != "" {
		collection := pr.database.Collection(pr.collection)
		objID, err := primitive.ObjectIDFromHex(id)
		if err != nil {
			return err
		}
		_, err = collection.UpdateByID(c, objID, bson.M{"$set": bson.M{"header": header}})
		if err != nil {
			return err
		}
	}
	awsClient := s3.New(pr.awsSession)
	bucket := os.Getenv("BUCKET_NAME")
	key := id + ".md"
	_, err := awsClient.DeleteObject(&s3.DeleteObjectInput{
		Bucket: aws.String(bucket),
		Key:    aws.String(key),
	})
	if err != nil {
		return err
	}

	uploader := s3manager.NewUploader(pr.awsSession)
	_, err = uploader.Upload(&s3manager.UploadInput{
		Bucket: aws.String(bucket),
		Key:    aws.String(key),
		Body:   file,
	})
	if err != nil {
		return err
	}
	return nil
}

func (pr *postRepository) GetOne(c *gin.Context, id string) (models.Post, error) {
	collection := pr.database.Collection(pr.collection)
	objID, err := primitive.ObjectIDFromHex(id)
	var post models.Post
	if err != nil {
		return post, err
	}
	err = collection.FindOne(c, bson.M{"_id": objID}).Decode(&post)
	if err != nil {
		return post, err
	}
	return post, nil
}

func (pr *postRepository) GetAll(c *gin.Context, page int, size int, categoryId string) ([]interfaces.GetAllPostResponse, int, error) {
	collection := pr.database.Collection(pr.collection)
	filter := bson.M{}
	if len(categoryId) > 0 {
		objId, err := primitive.ObjectIDFromHex(categoryId)
		if err != nil {
			return nil, 0, err
		}
		filter["category._id"] = objId
	}
	total, err := collection.CountDocuments(c, filter)
	if err != nil {
		return nil, 0, err
	}
	totalPages := int(math.Ceil(float64(total) / float64(size)))
	skip := (page - 1) * size

	cursor, err := collection.Find(
		c,
		filter,
		options.Find().SetSkip(int64(skip)).SetLimit(int64(size)),
	)
	if err != nil {
		return nil, 0, err
	}
	defer cursor.Close(c)
	var posts []interfaces.GetAllPostResponse
	for cursor.Next(c) {
		// var post interfaces.GetAllPostResponse
		var rawPost models.Post
		if err := cursor.Decode(&rawPost); err != nil {
			return nil, 0, err
		}
		post := interfaces.GetAllPostResponse{
			Id:       rawPost.Id,
			Header:   rawPost.Header,
			Date:     rawPost.Date,
			Category: rawPost.Category.Id,
		}
		posts = append(posts, post)
	}
	return posts, totalPages, nil
}
