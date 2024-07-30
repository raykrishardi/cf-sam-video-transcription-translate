package utils

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/transcribe"
)

func GetTranscribeClient(ctx context.Context, region string) (*transcribe.Client, error) {
	cfg, err := config.LoadDefaultConfig(ctx)
	if err != nil {
		return nil, err
	}

	transcribeClient := transcribe.NewFromConfig(cfg, func(o *transcribe.Options) {
		o.Region = region
	})

	return transcribeClient, nil
}
