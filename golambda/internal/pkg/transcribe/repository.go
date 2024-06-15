package transcribe

import "cf-sam-video-transcription-translate/internal/pkg/config"

var Repo *TranscribeRepository

type TranscribeRepository struct {
	App *config.AppConfig
}

func NewTranscribeRepository(app *config.AppConfig) *TranscribeRepository {
	return &TranscribeRepository{
		App: app,
	}
}

func NewTranscribe(repo *TranscribeRepository) {
	Repo = repo
}
