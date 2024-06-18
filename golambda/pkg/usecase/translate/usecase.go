package translate

import (
	"cf-sam-video-transcription-translate/config"
	"cf-sam-video-transcription-translate/pkg/repository"
)

type TranslateUseCase struct {
	AppConfig     *config.AppConfig
	TranslateRepo repository.TranslateRepo
}

func NewTranslateUseCase(appConfig *config.AppConfig, translateRepo repository.TranslateRepo) *TranslateUseCase {
	return &TranslateUseCase{
		AppConfig:     appConfig,
		TranslateRepo: translateRepo,
	}
}
