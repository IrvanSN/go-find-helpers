package thumbnail

import (
	"context"
	"fmt"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/credentials"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/google/uuid"
	"github.com/irvansn/go-find-helpers/entities"
	"log"
	"os"
)

type Thumbnail struct {
	ID           uuid.UUID `gorm:"type:varchar(100);"`
	JobID        uuid.UUID `gorm:"type:varchar(100);"`
	ImageKey     string    `gorm:"type:text;"`
	Description  string    `gorm:"type:varchar(100);"`
	PreSignedURL string    `gorm:"-"`
}

func FromUseCase(thumbnail *entities.Thumbnail) *Thumbnail {
	return &Thumbnail{
		ID:           thumbnail.ID,
		JobID:        thumbnail.JobID,
		ImageKey:     thumbnail.ImageKey,
		Description:  thumbnail.Description,
		PreSignedURL: thumbnail.PreSignedURL,
	}
}

func (t *Thumbnail) ToUseCase() *entities.Thumbnail {
	return &entities.Thumbnail{
		ID:           t.ID,
		JobID:        t.JobID,
		ImageKey:     t.ImageKey,
		Description:  t.Description,
		PreSignedURL: t.PreSignedURL,
	}
}

func (t *Thumbnail) GetPreSignedURL() error {
	var bucketName = os.Getenv("R2_BUCKET_NAME")
	var accountId = os.Getenv("R2_ACCOUNT_ID")
	var accessKeyId = os.Getenv("R2_ACCESS_KEY")
	var accessKeySecret = os.Getenv("R2_SECRET_KEY")
	var objectKey = t.ImageKey

	r2Resolver := aws.EndpointResolverWithOptionsFunc(func(service, region string, options ...interface{}) (aws.Endpoint, error) {
		return aws.Endpoint{
			URL: fmt.Sprintf("https://%s.r2.cloudflarestorage.com", accountId),
		}, nil
	})

	cfg, err := config.LoadDefaultConfig(context.TODO(),
		config.WithEndpointResolverWithOptions(r2Resolver),
		config.WithCredentialsProvider(credentials.NewStaticCredentialsProvider(accessKeyId, accessKeySecret, "")),
		config.WithRegion("auto"),
	)
	if err != nil {
		fmt.Println("Failed to create AWS configuration:", err)
		return err
	}

	client := s3.NewFromConfig(cfg)
	preSignClient := s3.NewPresignClient(client)
	request, err := preSignClient.PresignPutObject(context.TODO(), &s3.PutObjectInput{
		Bucket: aws.String(bucketName),
		Key:    aws.String(objectKey),
	})
	if err != nil {
		log.Printf("Couldn't get a presigned request to put %v:%v. Here's why: %v\n", bucketName, objectKey, err)
	}

	t.PreSignedURL = request.URL
	return nil
}
