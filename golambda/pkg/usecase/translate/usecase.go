package translate

import (
	"github.com/raykrishardi/cf-sam-video-transcription-translate/golambda/config"
	"github.com/raykrishardi/cf-sam-video-transcription-translate/golambda/pkg/repository"
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
