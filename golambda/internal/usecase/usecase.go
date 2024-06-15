package usecase

import (
	mcuc "cf-sam-video-transcription-translate/internal/usecase/mediaconvert"
	s3uc "cf-sam-video-transcription-translate/internal/usecase/s3"
	truc "cf-sam-video-transcription-translate/internal/usecase/transcribe"
	tluc "cf-sam-video-transcription-translate/internal/usecase/translate"
)

type UseCase struct {
	MediaConvertUseCase *mcuc.MediaConvertUseCase
	TranscribeUseCase   *truc.TranscribeUseCase
	TranslateUseCase    *tluc.TranslateUseCase
	S3UseCase           *s3uc.S3UseCase
}

func NewUseCase(mediaConvertUseCase *mcuc.MediaConvertUseCase, transcribeUseCase *truc.TranscribeUseCase, translateUseCase *tluc.TranslateUseCase, s3UseCase *s3uc.S3UseCase) *UseCase {
	return &UseCase{
		MediaConvertUseCase: mediaConvertUseCase,
		TranscribeUseCase:   transcribeUseCase,
		TranslateUseCase:    translateUseCase,
		S3UseCase:           s3UseCase,
	}
}
