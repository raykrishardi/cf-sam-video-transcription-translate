package translate

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/translate"
)

func (repo *TranslateTranscriptionRepository) GetTranslateClient(ctx context.Context) (*translate.Client, error) {
	cfg, err := config.LoadDefaultConfig(ctx)
	if err != nil {
		return nil, err
	}

	translateClient := translate.NewFromConfig(cfg)

	return translateClient, nil
}
