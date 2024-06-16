package usecase

import (
	mcuc "cf-sam-video-transcription-translate/pkg/usecase/mediaconvert"
	osuc "cf-sam-video-transcription-translate/pkg/usecase/objectstore"
	truc "cf-sam-video-transcription-translate/pkg/usecase/transcribe"
	tluc "cf-sam-video-transcription-translate/pkg/usecase/translate"
)

type UseCase struct {
	MediaConvertUseCase *mcuc.MediaConvertUseCase
	TranscribeUseCase   *truc.TranscribeUseCase
	TranslateUseCase    *tluc.TranslateUseCase
	ObjectStoreUseCase  *osuc.ObjectStoreUseCase
}

func NewUseCase(mediaConvertUseCase *mcuc.MediaConvertUseCase, transcribeUseCase *truc.TranscribeUseCase, translateUseCase *tluc.TranslateUseCase, objectStoreUseCase *osuc.ObjectStoreUseCase) *UseCase {
	return &UseCase{
		MediaConvertUseCase: mediaConvertUseCase,
		TranscribeUseCase:   transcribeUseCase,
		TranslateUseCase:    translateUseCase,
		ObjectStoreUseCase:  objectStoreUseCase,
	}
}
