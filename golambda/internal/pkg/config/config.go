package config

type AppConfig struct {
	AWSRegion               string
	VideoBucketName         *string
	AudioBucketName         *string
	MediaConvertIamRoleArn  *string
	AWSMediaConvertEndpoint *string
}
