package entity

type ConvertMP4ToMP3Input struct {
	Role           string
	OutMP3Settings MP3Settings
	OutS3Uri       string
	InS3Uri        string
}

type MP3Settings struct {
	RateControlMode string `json:"rate_control_mode"`
	VbrQuality      *int32 `json:"vbr_quality,omitempty"`
	BitRate         *int32 `json:"bit_rate,omitempty"`
}

type GetJobInput struct {
	ID string `json:"id"`
}

type GetJobOutput struct {
	Status string `json:"status"`
}
