package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"os"

	"github.com/aws/aws-lambda-go/lambda"

	"cf-sam-video-transcription-translate/config"
	trrepo "cf-sam-video-transcription-translate/pkg/repository/transcribe"
	truc "cf-sam-video-transcription-translate/pkg/usecase/transcribe"

	"cf-sam-video-transcription-translate/pkg/entity"
	"cf-sam-video-transcription-translate/pkg/utility"
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

	// Initialise repositories
	trRepo := trrepo.NewTranscribeRepository(appConfig)
	trrepo.NewTranscribe(trRepo)

	// Initialise specific usecases
	trUC := truc.NewTranscribeUseCase(ctx, trRepo)

	// Initialise global usecase (if necessary)
	// uc := usecase.NewUseCase(nil, trUC)

	// Business logic
	autoLanguageDetection := true
	inBucketDirPath := utility.Split(event.Detail.Object.Key, "/", true, false)
	inBucketFileName := utility.Split(event.Detail.Object.Key, "/", false, true) // file name with extension (e.g. hello.mp3)
	inBucketFileNameWithoutExtension := utility.GetFileNameOrExtension(inBucketFileName, false)
	outBucketObjectKey := fmt.Sprintf("%s/%s", inBucketDirPath, inBucketFileNameWithoutExtension)
	transcribeMP3ToSRTInput := entity.TranscribeMP3ToSRTInput{
		OutBucketName:      appConfig.TranscriptionBucketName,
		OutBucketObjectKey: &outBucketObjectKey,
		InS3Uri:            fmt.Sprintf("s3://%s/%s", event.Detail.Bucket.Name, event.Detail.Object.Key),
		InFileName:         inBucketFileName,
		IdentifyLanguage:   &autoLanguageDetection,
	}
	transcribeMP3ToSRTOutput, err := trUC.TranscribeMP3ToSRT(ctx, transcribeMP3ToSRTInput)
	if err != nil {
		log.Fatalf("Unable to transcribe mp3 from %s bucket: %v\n", appConfig.AudioBucketName, err)
	}

	resultBytes, err := json.Marshal(transcribeMP3ToSRTOutput)
	if err != nil {
		log.Fatalf("Error serializing transcribeMP3ToSRTOutput to JSON:%v\n", err)
	}
	log.Printf("result: %s\n", resultBytes)

	return resultBytes, nil
}

func main() {
	lambda.Start(handler)
}
