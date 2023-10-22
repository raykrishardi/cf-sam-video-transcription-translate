package usecase

import (
	s3uc "cf-sam-video-transcription-translate/internal/usecase/s3"
)

type UseCase struct {
	S3UseCase *s3uc.S3UseCase
}

func NewUseCase(s3UseCase *s3uc.S3UseCase) *UseCase {
	return &UseCase{
		S3UseCase: s3UseCase,
	}
}
