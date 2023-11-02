package usecase

import (
	mcuc "cf-sam-video-transcription-translate/internal/usecase/mediaconvert"
	truc "cf-sam-video-transcription-translate/internal/usecase/transcribe"
)

type UseCase struct {
	MediaConvertUseCase *mcuc.MediaConvertUseCase
	TranscribeUseCase   *truc.TranscribeUseCase
}

func NewUseCase(mediaConvertUseCase *mcuc.MediaConvertUseCase, transcribeUseCase *truc.TranscribeUseCase) *UseCase {
	return &UseCase{
		MediaConvertUseCase: mediaConvertUseCase,
		TranscribeUseCase:   transcribeUseCase,
	}
}
