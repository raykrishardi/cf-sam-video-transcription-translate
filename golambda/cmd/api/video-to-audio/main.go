package main

import (
	"cf-sam-video-transcription-translate/internal/entity/eventbridge"
	"context"
	"encoding/json"
	"fmt"
	"log"
	"os"

	"github.com/aws/aws-lambda-go/lambda"

	"cf-sam-video-transcription-translate/internal/pkg/config"
	"cf-sam-video-transcription-translate/internal/pkg/helper"
	mcrepo "cf-sam-video-transcription-translate/internal/pkg/mediaconvert"
	mcuc "cf-sam-video-transcription-translate/internal/usecase/mediaconvert"
)

var (
	AWS_REGION                         = os.Getenv("AWS_REGION")
	SOURCE_BUCKET_NAME                 = os.Getenv("SOURCE_BUCKET_NAME")
	DESTINATION_BUCKET_NAME            = os.Getenv("DESTINATION_BUCKET_NAME")
	MEDIA_CONVERT_DEFAULT_IAM_ROLE_ARN = os.Getenv("MEDIA_CONVERT_DEFAULT_IAM_ROLE_ARN")
	AWS_MEDIA_CONVERT_ENDPOINT         = os.Getenv("AWS_MEDIA_CONVERT_ENDPOINT")
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
		VideoBucketName:         &SOURCE_BUCKET_NAME,
		AudioBucketName:         &DESTINATION_BUCKET_NAME,
		MediaConvertIamRoleArn:  &MEDIA_CONVERT_DEFAULT_IAM_ROLE_ARN,
		AWSMediaConvertEndpoint: &AWS_MEDIA_CONVERT_ENDPOINT,
	}

	// Initialise repositories
	mcRepo := mcrepo.NewMediaConvertRepository(appConfig)
	mcrepo.NewMediaConvert(mcRepo)

	// Initialise specific usecases
	mcUC := mcuc.NewMediaConvertUseCase(ctx, mcRepo)

	// Initialise global usecase (if necessary)
	// uc := usecase.NewUseCase(mcUC, nil)

	// Business logic
	bitRate := int32(192000)
	convertMP4ToMP3Input := mcuc.ConvertMP4ToMP3Input{
		Role:     *appConfig.MediaConvertIamRoleArn,
		InS3Uri:  fmt.Sprintf("s3://%s/%s", event.Detail.Bucket.Name, event.Detail.Object.Key),
		OutS3Uri: fmt.Sprintf("s3://%s/%s/", *appConfig.AudioBucketName, helper.Split(event.Detail.Object.Key, "/", true, false)),
		OutMP3Settings: mcuc.MP3Settings{
			RateControlMode: "CBR",
			BitRate:         &bitRate,
		},
	}
	convertMP4ToMP3Output, err := mcUC.ConvertMP4ToMP3(ctx, convertMP4ToMP3Input)
	if err != nil {
		log.Fatalf("Unable to convert mp4 to mp3 for %s bucket: %v\n", *appConfig.VideoBucketName, err)
	}

	resultBytes, err := json.Marshal(convertMP4ToMP3Output)
	if err != nil {
		log.Fatalf("Error serializing convertMP4ToMP3Output to JSON:%v\n", err)
	}
	log.Printf("result: %s\n", resultBytes)

	return resultBytes, nil
}

func main() {
	lambda.Start(handler)
}
