package usecase

import (
	mcuc "github.com/raykrishardi/cf-sam-video-transcription-translate/pkg/usecase/mediaconvert"
	osuc "github.com/raykrishardi/cf-sam-video-transcription-translate/pkg/usecase/objectstore"
	truc "github.com/raykrishardi/cf-sam-video-transcription-translate/pkg/usecase/transcribe"
	tluc "github.com/raykrishardi/cf-sam-video-transcription-translate/pkg/usecase/translate"
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
