package translate

import "cf-sam-video-transcription-translate/config"

var Repo *TranslateTranscriptionRepository

type TranslateTranscriptionRepository struct {
	App *config.AppConfig
}

func NewTranslateTranscriptionRepository(app *config.AppConfig) *TranslateTranscriptionRepository {
	return &TranslateTranscriptionRepository{
		App: app,
	}
}

func NewTranslate(repo *TranslateTranscriptionRepository) {
	Repo = repo
}
