package main

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/aws/aws-lambda-go/lambda"
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

func handler(context context.Context, event interface{}) ([]byte, error) {
	eventByte, _ := json.Marshal(event)
	eventStr := string(eventByte)
	fmt.Println(eventStr)

	return eventByte, nil
}

func main() {
	lambda.Start(handler)
}
