package repository

import (
	"cf-sam-video-transcription-translate/pkg/entity"
	"context"
)

type MediaConvertRepo interface {
	ConvertMP4ToMP3(ctx context.Context, params entity.ConvertMP4ToMP3Input) error
}
