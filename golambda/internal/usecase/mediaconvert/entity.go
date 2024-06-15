package mediaconvert

type ConvertMP4ToMP3Input struct {
	Role           string
	OutMP3Settings MP3Settings
	OutS3Uri       string
	InS3Uri        string
}

type MP3Settings struct {
	RateControlMode string
	VbrQuality      *int32
	BitRate         *int32
}
