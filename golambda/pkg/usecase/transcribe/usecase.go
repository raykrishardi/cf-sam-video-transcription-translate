package transcribe

import (
	"cf-sam-video-transcription-translate/config"
	"cf-sam-video-transcription-translate/pkg/repository"
)

type TranscribeUseCase struct {
	AppConfig      *config.AppConfig
	TranscribeRepo repository.TranscribeRepo
}

func NewTranscribeUseCase(appConfig *config.AppConfig, transcribeRepo repository.TranscribeRepo) *TranscribeUseCase {
	return &TranscribeUseCase{
		AppConfig:      appConfig,
		TranscribeRepo: transcribeRepo,
	}
}
