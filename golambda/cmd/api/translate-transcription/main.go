package main

import (
	"cf-sam-video-transcription-translate/internal/entity/eventbridge"
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"

	"github.com/aws/aws-lambda-go/lambda"

	"cf-sam-video-transcription-translate/internal/pkg/config"
	s3repo "cf-sam-video-transcription-translate/internal/pkg/s3"
	tlrepo "cf-sam-video-transcription-translate/internal/pkg/translate"
	s3uc "cf-sam-video-transcription-translate/internal/usecase/s3"
	tluc "cf-sam-video-transcription-translate/internal/usecase/translate"
)

var (
	AWS_REGION              = os.Getenv("AWS_REGION")
	SOURCE_BUCKET_NAME      = os.Getenv("SOURCE_BUCKET_NAME")
	DESTINATION_BUCKET_NAME = os.Getenv("DESTINATION_BUCKET_NAME")
)

func handler(ctx context.Context, event eventbridge.S3) ([]byte, error) {
	eventBytes, err := json.Marshal(event)
	if err != nil {
		log.Fatalf("Error serializing event to JSON:%v\n", err)
	}
	log.Printf("event: %s\n", eventBytes)

	// Initialise app config
	appConfig := &config.AppConfig{
		AWSRegion:               AWS_REGION,
		TranscriptionBucketName: &SOURCE_BUCKET_NAME,
		TranslationBucketName:   &DESTINATION_BUCKET_NAME,
	}

	// Initialise repositories
	s3Repo := s3repo.NewS3Repository(appConfig)
	s3repo.NewS3(s3Repo)
	tlRepo := tlrepo.NewTranslateTranscriptionRepository(appConfig)
	tlrepo.NewTranslate(tlRepo)

	// Initialise specific usecases
	s3UC := s3uc.NewS3UseCase(ctx, s3Repo)
	tlUC := tluc.NewTranslateUseCase(ctx, tlRepo)

	// Initialise global usecase (if necessary)
	// uc := usecase.NewUseCase(nil, nil, tlUC, s3UC)

	// Business logic
	s3GetObjectInput := s3uc.GetObjectInput{
		BucketName: event.Detail.Bucket.Name,
		Key:        event.Detail.Object.Key,
	}
	s3GetObjectOutput, err := s3UC.GetObject(ctx, s3GetObjectInput)
	if err != nil {
		log.Fatalf("Unable to get s3 content for key %s and bucket %s: %v\n", s3GetObjectInput.Key, s3GetObjectInput.BucketName, err)
	}

	s3ObjectBytes, err := ioutil.ReadAll(s3GetObjectOutput.Body)
	if err != nil {
		log.Fatalf("Error reading s3 content for key %s and bucket %s: %v\n", s3GetObjectInput.Key, s3GetObjectInput.BucketName, err)
	}

	// List of supported language code (https://docs.aws.amazon.com/translate/latest/dg/what-is-languages.html)
	sourceLanguageCode := "auto"
	targetLanguageCode := "id"
	translateDocumentInput := tluc.TranslateDocumentInput{
		Content:            s3ObjectBytes,
		ContentType:        "text/plain",
		SourceLanguageCode: &sourceLanguageCode,
		TargetLanguageCode: &targetLanguageCode,
	}
	translateDocumentOutput, err := tlUC.TranslateDocument(ctx, translateDocumentInput)
	if err != nil {
		log.Fatalf("Unable to translate s3 content for key %s and bucket %s: %v\n", s3GetObjectInput.Key, s3GetObjectInput.BucketName, err)
	}

	s3PutObjectInput := s3uc.PutObjectInput{
		BucketName: *s3UC.S3Repo.App.TranslationBucketName,
		Key:        fmt.Sprintf("%s/%s", *translateDocumentInput.TargetLanguageCode, event.Detail.Object.Key),
		Body:       translateDocumentOutput.TranslatedDocument.Content,
	}
	_, err = s3UC.PutObject(ctx, s3PutObjectInput)
	if err != nil {
		log.Fatalf("Unable to upload translated content to key %s and bucket %s: %v\n", s3PutObjectInput.Key, s3PutObjectInput.BucketName, err)
	}

	resultBytes, err := json.Marshal(translateDocumentOutput)
	if err != nil {
		log.Fatalf("Error serializing translateDocumentOutput to JSON:%v\n", err)
	}
	log.Printf("result: %s\n", resultBytes)

	return resultBytes, nil
}

func main() {
	lambda.Start(handler)
}
