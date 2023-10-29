package usecase

import (
	mcuc "cf-sam-video-transcription-translate/internal/usecase/mediaconvert"
)

type UseCase struct {
	MediaConvertUseCase *mcuc.MediaConvertUseCase
}

func NewUseCase(mediaConvertUseCase *mcuc.MediaConvertUseCase) *UseCase {
	return &UseCase{
		MediaConvertUseCase: mediaConvertUseCase,
	}
}
