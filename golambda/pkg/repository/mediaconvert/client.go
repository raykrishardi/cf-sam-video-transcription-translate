package mediaconvert

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/mediaconvert"
)

func (repo *MediaConvertRepository) GetMediaConvertClient(ctx context.Context) (*mediaconvert.Client, error) {
	cfg, err := config.LoadDefaultConfig(ctx)
	if err != nil {
		return nil, err
	}

	mediaConvertClient := mediaconvert.NewFromConfig(cfg, func(o *mediaconvert.Options) {
		o.BaseEndpoint = &repo.App.AWSMediaConvertEndpoint
	})

	return mediaConvertClient, nil
}
