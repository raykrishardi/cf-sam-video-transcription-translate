package repository

import (
	"context"

	"github.com/raykrishardi/cf-sam-video-transcription-translate/pkg/entity"
)

type MediaConvertRepo interface {
	ConvertMP4ToMP3(ctx context.Context, params entity.ConvertMP4ToMP3Input) (string, error)
}

type ObjectStoreRepo interface {
	GetObject(ctx context.Context, params entity.GetObjectInput) ([]byte, error)
	PutObject(ctx context.Context, params entity.PutObjectInput) error
}

type TranscribeRepo interface {
	TranscribeMP3ToSRT(ctx context.Context, params entity.TranscribeMP3ToSRTInput) error
}

type TranslateRepo interface {
	TranslateDocument(ctx context.Context, params entity.TranslateDocumentInput) ([]byte, error)
}
