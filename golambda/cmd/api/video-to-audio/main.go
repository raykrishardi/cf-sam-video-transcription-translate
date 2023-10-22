package main

import (
	"cf-sam-video-transcription-translate/internal/entity/eventbridge"
	"cf-sam-video-transcription-translate/internal/usecase"
	"context"
	"encoding/json"
	"log"
	"os"

	"github.com/aws/aws-lambda-go/lambda"

	"cf-sam-video-transcription-translate/internal/pkg/config"
	s3repo "cf-sam-video-transcription-translate/internal/pkg/s3"
	s3uc "cf-sam-video-transcription-translate/internal/usecase/s3"
)

var (
	AWS_REGION         = os.Getenv("AWS_REGION")
	SOURCE_BUCKET_NAME = os.Getenv("SOURCE_BUCKET_NAME")
)

func handler(ctx context.Context, event eventbridge.S3) ([]byte, error) {
	eventBytes, _ := json.Marshal(event)
	log.Printf("event: %s\n", eventBytes)

	// Initialise app config
	appConfig := &config.AppConfig{
		AWSRegion:       AWS_REGION,
		VideoBucketName: SOURCE_BUCKET_NAME,
	}

	// Initialise repositories
	s3Repo := s3repo.NewS3Repository(appConfig)
	s3repo.NewS3(s3Repo)

	// Initialise specific usecases
	s3UC := s3uc.NewS3UseCase(s3Repo)

	// Initialise global usecase
	uc := usecase.NewUseCase(s3UC)

	// Business logic
	result, err := uc.S3UseCase.ListBucket(ctx, appConfig.VideoBucketName)
	if err != nil {
		log.Fatalf("Unable to list %s bucket content: %v", appConfig.VideoBucketName, err)
	}

	for _, object := range result.Contents {
		log.Printf("key=%s size=%d", *object.Key, object.Size)
	}

	return eventBytes, nil
}

func main() {
	lambda.Start(handler)
}
