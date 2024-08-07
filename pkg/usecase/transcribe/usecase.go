package transcribe

import (
	"github.com/raykrishardi/cf-sam-video-transcription-translate/config"
	"github.com/raykrishardi/cf-sam-video-transcription-translate/pkg/repository"
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
