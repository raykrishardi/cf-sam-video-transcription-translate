package objectstore

import (
	"cf-sam-video-transcription-translate/config"
	"cf-sam-video-transcription-translate/pkg/repository"
	"context"

	"cf-sam-video-transcription-translate/pkg/entity"
)

type ObjectStoreUseCase struct {
	AppConfig       *config.AppConfig
	ObjectStoreRepo repository.ObjectStoreRepo
}

func NewObjectStoreUseCase(appConfig *config.AppConfig, objectStoreRepo repository.ObjectStoreRepo) *ObjectStoreUseCase {
	return &ObjectStoreUseCase{
		AppConfig:       appConfig,
		ObjectStoreRepo: objectStoreRepo,
	}
}

func (uc *ObjectStoreUseCase) GetObject(ctx context.Context, params entity.GetObjectInput) ([]byte, error) {
	return uc.ObjectStoreRepo.GetObject(ctx, params)
}

func (uc *ObjectStoreUseCase) PutObject(ctx context.Context, params entity.PutObjectInput) error {
	return uc.ObjectStoreRepo.PutObject(ctx, params)
}
