package repository

import (
	"errors"
	"io/ioutil"
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
