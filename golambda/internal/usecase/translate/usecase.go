package translate

import (
	tlrepo "cf-sam-video-transcription-translate/internal/pkg/translate"
	"context"

	"github.com/aws/aws-sdk-go-v2/service/translate"
	"github.com/aws/aws-sdk-go-v2/service/translate/types"
)

type TranslateUseCase struct {
	TranslateRepo *tlrepo.TranslateTranscriptionRepository
	Client        *translate.Client
}

func NewTranslateUseCase(ctx context.Context, tlRepo *tlrepo.TranslateTranscriptionRepository) *TranslateUseCase {
	uc := &TranslateUseCase{}

	client, err := tlRepo.GetTranslateClient(ctx)
	if err != nil {
		return uc
	}

	uc.TranslateRepo = tlRepo
	uc.Client = client

	return uc
}

func (uc *TranslateUseCase) TranslateDocument(ctx context.Context, params TranslateDocumentInput) (*translate.TranslateDocumentOutput, error) {
	tdi := &translate.TranslateDocumentInput{
		Document: &types.Document{
			Content:     params.Content,
			ContentType: &params.ContentType,
		},
		SourceLanguageCode: params.SourceLanguageCode,
		TargetLanguageCode: params.TargetLanguageCode,
	}

	return uc.Client.TranslateDocument(ctx, tdi)
}
