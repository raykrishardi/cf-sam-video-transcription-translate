package repository

import (
	"cf-sam-video-transcription-translate/pkg/entity"
	"context"
)

type MediaConvertRepo interface {
	ConvertMP4ToMP3(ctx context.Context, params entity.ConvertMP4ToMP3Input) error
}

type ObjectStoreRepo interface {
	GetObject(ctx context.Context, params entity.GetObjectInput) ([]byte, error)
	PutObject(ctx context.Context, params entity.PutObjectInput) error
}

type TranscribeRepo interface {
	TranscribeMP3ToSRT(ctx context.Context, params entity.TranscribeMP3ToSRTInput) error
}
