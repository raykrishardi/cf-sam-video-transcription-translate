package translate

import (
	"cf-sam-video-transcription-translate/config"
	"cf-sam-video-transcription-translate/pkg/entity"
	"context"

	"github.com/aws/aws-sdk-go-v2/service/translate"
	"github.com/aws/aws-sdk-go-v2/service/translate/types"
)

type AWSTranslateRepo struct {
	App    *config.AppConfig
	Client *translate.Client
}

func NewAWSTranslateRepo(app *config.AppConfig, client *translate.Client) *AWSTranslateRepo {
	return &AWSTranslateRepo{
		App:    app,
		Client: client,
	}
}

func (r *AWSTranslateRepo) TranslateDocument(ctx context.Context, params entity.TranslateDocumentInput) ([]byte, error) {
	tdi := &translate.TranslateDocumentInput{
		Document: &types.Document{
			Content:     params.Content,
			ContentType: &params.ContentType,
		},
		SourceLanguageCode: params.SourceLanguageCode,
		TargetLanguageCode: params.TargetLanguageCode,
	}

	tdo, err := r.Client.TranslateDocument(ctx, tdi)
	if err != nil {
		return nil, err
	}

	translatedContent := tdo.TranslatedDocument.Content
	return translatedContent, nil
}
