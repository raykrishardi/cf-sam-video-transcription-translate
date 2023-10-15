package main

import (
	"cf-sam-video-transcription-translate/api/model/aws/eventbridge"
	"context"
	"encoding/json"
	"go/types"
	"log"
	"os"

	"github.com/aws/aws-lambda-go/lambda"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/s3"
)

var (
	AWS_REGION         = os.Getenv("AWS_REGION")
	SOURCE_BUCKET_NAME = os.Getenv("SOURCE_BUCKET_NAME")
)

/*
Description: This lambda handler reads an EC2 Instance Change Notification from the inputted AWSEvent, then updates the

	Detail Type (description of the type of AWS event).

Parameters:
  - context: Object that provides information about the Lambda invocation, function, and execution environment.
  - awsEvent: AWS EventBridge object that provides information such as the source, region, account the
    event is associated with.

Output: Stream ([]byte) of the updated AWSEvent after function logic completed.
*/
// func handler(context context.Context, awsEvent ec2.AWSEvent) ([]byte, error) {
// 	// Retrieve the ec2 notification from the event
// 	ec2InstanceStateChangeNotification := awsEvent.Detail

// 	// Developers write your event-driven business logic code here!
// 	fmt.Println("Instance " + ec2InstanceStateChangeNotification.InstanceId + " transitioned to " + ec2InstanceStateChangeNotification.State)

// 	// Make updates to the event payload
// 	awsEvent.SetDetailType("HelloWorldFunction updated event of " + awsEvent.DetailType)

// 	// Return event as stream for further processing
// 	return marshaller.Marshal(awsEvent)

// }

func handler(ctx context.Context, event eventbridge.S3) ([]byte, error) {
	eventBytes, _ := json.Marshal(event)
	log.Println(string(eventBytes))

	cfg, err := config.LoadDefaultConfig(ctx)
	if err != nil {
		log.Printf("failed to load default config: %s", err)
		return eventBytes, err
	}

	s3Client := s3.NewFromConfig(cfg)

	result, err := s3Client.ListObjectsV2(ctx, &s3.ListObjectsV2Input{
		Bucket: &event.Detail.Bucket.Name,
	})
	var contents []types.Object
	if err != nil {
		log.Printf("Couldn't list objects in bucket %v. Here's why: %v\n", bucketName, err)
	}

	log.Println(contents)

	// for _, record := range event.Records {
	// 	bucket := record.S3.Bucket.Name
	// 	key := record.S3.Object.URLDecodedKey
	// 	headOutput, err := s3Client.HeadObject(ctx, &s3.HeadObjectInput{
	// 		Bucket: &bucket,
	// 		Key:    &key,
	// 	})
	// 	if err != nil {
	// 		log.Printf("error getting head of object %s/%s: %s", bucket, key, err)
	// 		return eventBytes, err
	// 	}
	// 	log.Printf("successfully retrieved %s/%s of type %s", bucket, key, *headOutput.ContentType)
	// }

	return eventBytes, nil
}

func main() {
	lambda.Start(handler)
}
