package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"os"

	"github.com/aws/aws-lambda-go/lambda"

	"cf-sam-video-transcription-translate/config"
	mcrepo "cf-sam-video-transcription-translate/pkg/repository/mediaconvert"
	mcuc "cf-sam-video-transcription-translate/pkg/usecase/mediaconvert"

	"cf-sam-video-transcription-translate/pkg/entity"
	"cf-sam-video-transcription-translate/pkg/utils"
)

var (
	AWS_REGION                         = os.Getenv("AWS_REGION")
	SOURCE_BUCKET_NAME                 = os.Getenv("SOURCE_BUCKET_NAME")
	DESTINATION_BUCKET_NAME            = os.Getenv("DESTINATION_BUCKET_NAME")
	MEDIA_CONVERT_DEFAULT_IAM_ROLE_ARN = os.Getenv("MEDIA_CONVERT_DEFAULT_IAM_ROLE_ARN")
	AWS_MEDIA_CONVERT_ENDPOINT         = os.Getenv("AWS_MEDIA_CONVERT_ENDPOINT")
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
		VideoBucketName:         SOURCE_BUCKET_NAME,
		AudioBucketName:         DESTINATION_BUCKET_NAME,
		MediaConvertIamRoleArn:  MEDIA_CONVERT_DEFAULT_IAM_ROLE_ARN,
		AWSMediaConvertEndpoint: AWS_MEDIA_CONVERT_ENDPOINT,
	}

	// Initialise AWS MediaConvert client
	awsMediaConvertClient, err := utils.GetAWSMediaConvertClient(ctx, appConfig.AWSMediaConvertEndpoint)
	if err != nil {
		log.Fatalf("Error getting AWS MediaConvert client:%v\n", err)
	}

	// Initialise repositories
	awsMediaConvertRepo := mcrepo.NewAWSMediaConvertRepo(appConfig, awsMediaConvertClient)

	// Initialise specific usecases
	mcUC := mcuc.NewMediaConvertUseCase(appConfig, awsMediaConvertRepo)

	// Initialise global usecase (if necessary)
	// uc := usecase.NewUseCase(mcUC, nil)

	// Business logic
	bitRate := int32(192000)
	convertMP4ToMP3Input := entity.ConvertMP4ToMP3Input{
		Role:     appConfig.MediaConvertIamRoleArn,
		InS3Uri:  fmt.Sprintf("s3://%s/%s", event.Detail.Bucket.Name, event.Detail.Object.Key),
		OutS3Uri: fmt.Sprintf("s3://%s/%s/", appConfig.AudioBucketName, utils.GetDirPathOrFileName(event.Detail.Object.Key, "/", true, false)),
		OutMP3Settings: entity.MP3Settings{
			RateControlMode: "CBR",
			BitRate:         &bitRate,
		},
	}
	err = mcUC.ConvertMP4ToMP3(ctx, convertMP4ToMP3Input)
	if err != nil {
		log.Fatalf("Unable to convert mp4 to mp3 for %s bucket: %v\n", appConfig.VideoBucketName, err)
	}

	response := fmt.Sprintf("Successfully converted  %s from mp4 to mp3", event.Detail.Object.Key)
	log.Println(response)
	return []byte(response), nil
}

func main() {
	lambda.Start(handler)
}
