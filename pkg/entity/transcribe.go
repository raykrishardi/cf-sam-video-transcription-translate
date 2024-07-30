package entity

type TranscribeMP3ToSRTInput struct {
	OutBucketName      string
	OutBucketObjectKey *string
	InS3Uri            string
	InFileName         string
	IdentifyLanguage   *bool
}
