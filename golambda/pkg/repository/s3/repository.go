package s3

import "cf-sam-video-transcription-translate/config"

var Repo *S3Repository

type S3Repository struct {
	App *config.AppConfig
}

func NewS3Repository(app *config.AppConfig) *S3Repository {
	return &S3Repository{
		App: app,
	}
}

func NewS3(repo *S3Repository) {
	Repo = repo
}
