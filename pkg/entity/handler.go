package entity

type ConvertMP4ToMP3RequestPayload struct {
	SourceBucketName      string      `json:"source_bucket_name"`
	SourceKey             string      `json:"source_key"`
	DestinationBucketName string      `json:"destination_bucket_name"`
	DestinationKey        string      `json:"destination_key"`
	MP3Settings           MP3Settings `json:"mp3_settings"`
}
