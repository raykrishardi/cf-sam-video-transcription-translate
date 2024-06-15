package transcribe

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/transcribe"
)

func (repo *TranscribeRepository) GetTranscribeClient(ctx context.Context) (*transcribe.Client, error) {
	cfg, err := config.LoadDefaultConfig(ctx)
	if err != nil {
		return nil, err
	}

	transcribeClient := transcribe.NewFromConfig(cfg, func(o *transcribe.Options) {
		o.Region = repo.App.AWSRegion
	})

	return transcribeClient, nil
}
