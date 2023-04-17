package repository

import (
	"errors"
	"mime/multipart"
	"os"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
	"github.com/gin-gonic/gin"
	"github.com/slyxh2/golang-blog/interfaces"
	"github.com/slyxh2/golang-blog/models"
	"go.mongodb.org/mongo-driver/mongo"
)

type postRepository struct {
	database   *mongo.Database
	collection string
}

func NewPostRepository(db *mongo.Database) *postRepository {
	return &postRepository{
		database:   db,
		collection: interfaces.CollectionPost,
	}
}

func (pr *postRepository) Upload(c *gin.Context, file multipart.File, post *models.Post) (*s3manager.UploadOutput, error) {
	if file == nil {
		return nil, errors.New("file is nil")
	}

	cred := credentials.NewStaticCredentials(os.Getenv("BUCKET_KEY_ID"), os.Getenv("BUCKET_SECRET_KEY"), "")
	// The session the S3 Uploader will use
	sess, err := session.NewSession(&aws.Config{
		Region:      aws.String("us-east-1"),
		Credentials: cred,
	})
	if err != nil {
		return nil, err
	}
	// Create an uploader with the session and default options
	uploader := s3manager.NewUploader(sess)
	bucket := os.Getenv("BUCKET_NAME")
	key := post.Header + ".md"

	// Upload the file to S3.
	result, err := uploader.Upload(&s3manager.UploadInput{
		Bucket: aws.String(bucket),
		Key:    aws.String(key),
		Body:   file,
	})
	if err != nil {
		return nil, err
	}

	collection := pr.database.Collection(pr.collection)
	_, err = collection.InsertOne(c, post)
	return result, err
}
