package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"os"

	"github.com/aws/aws-lambda-go/lambda"

	truc "github.com/raykrishardi/cf-sam-video-transcription-translate/golambda/pkg/usecase/transcribe"

	trrepo "github.com/raykrishardi/cf-sam-video-transcription-translate/golambda/pkg/repository/transcribe"

	"github.com/raykrishardi/cf-sam-video-transcription-translate/golambda/config"

	"github.com/raykrishardi/cf-sam-video-transcription-translate/golambda/pkg/utils"

	"github.com/raykrishardi/cf-sam-video-transcription-translate/golambda/pkg/entity"
)

var (
	AWS_REGION              = os.Getenv("AWS_REGION")
	SOURCE_BUCKET_NAME      = os.Getenv("SOURCE_BUCKET_NAME")
	DESTINATION_BUCKET_NAME = os.Getenv("DESTINATION_BUCKET_NAME")
)

func handler(ctx context.Context, event entity.AWSEventBridgeS3Event) ([]byte, error) {
	eventBytes, err := json.Marshal(event)
	if err != nil {
		log.Fatalf("Error serializing event to JSON:%v\n", err)
	}
	log.Printf("event: %s\n", eventBytes)

	// Initialise app config
	appConfig := &config.AppConfig{
		AWSRegion:               AWS_REGION,
		AudioBucketName:         SOURCE_BUCKET_NAME,
		TranscriptionBucketName: DESTINATION_BUCKET_NAME,
	}

	// Initialise Transcribe client
	awsTranscribeClient, err := utils.GetTranscribeClient(ctx, appConfig.AWSRegion)
	if err != nil {
		log.Fatalf("Error getting transcribe client:%v\n", err)
	}

	// Initialise repositories
	awsTranscribeRepo := trrepo.NewAWSTranscribeRepo(appConfig, awsTranscribeClient)

	// Initialise specific usecases
	trUC := truc.NewTranscribeUseCase(appConfig, awsTranscribeRepo)

	// Initialise global usecase (if necessary)
	// uc := usecase.NewUseCase(nil, trUC)

	// Business logic
	autoLanguageDetection := true
	inBucketDirPath := utils.GetDirPathOrFileName(event.Detail.Object.Key, "/", true, false)
	inBucketFileName := utils.GetDirPathOrFileName(event.Detail.Object.Key, "/", false, true) // file name with extension (e.g. hello.mp3)
	inBucketFileNameWithoutExtension := utils.GetFileNameOrExtension(inBucketFileName, false)
	outBucketObjectKey := fmt.Sprintf("%s/%s", inBucketDirPath, inBucketFileNameWithoutExtension)
	transcribeMP3ToSRTInput := entity.TranscribeMP3ToSRTInput{
		OutBucketName:      appConfig.TranscriptionBucketName,
		OutBucketObjectKey: &outBucketObjectKey,
		InS3Uri:            fmt.Sprintf("s3://%s/%s", event.Detail.Bucket.Name, event.Detail.Object.Key),
		InFileName:         inBucketFileName,
		IdentifyLanguage:   &autoLanguageDetection,
	}
	err = trUC.TranscribeRepo.TranscribeMP3ToSRT(ctx, transcribeMP3ToSRTInput)
	if err != nil {
		log.Fatalf("Unable to transcribe mp3 from %s bucket: %v\n", appConfig.AudioBucketName, err)
	}

	response := fmt.Sprintf("Successfully transcribed %s mp3 file", inBucketFileName)
	log.Println(response)
	return []byte(response), nil
}

func main() {
	lambda.Start(handler)
}
