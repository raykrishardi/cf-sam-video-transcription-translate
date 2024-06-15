package mediaconvert

import "cf-sam-video-transcription-translate/config"

var Repo *MediaConvertRepository

type MediaConvertRepository struct {
	App *config.AppConfig
}

func NewMediaConvertRepository(app *config.AppConfig) *MediaConvertRepository {
	return &MediaConvertRepository{
		App: app,
	}
}

func NewMediaConvert(repo *MediaConvertRepository) {
	Repo = repo
}
