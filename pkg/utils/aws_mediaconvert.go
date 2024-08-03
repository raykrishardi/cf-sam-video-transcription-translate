package utils

import (
	"context"
	"fmt"

	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/mediaconvert"
)

func GetAWSMediaConvertClient(ctx context.Context, endpoint string) (*mediaconvert.Client, error) {
	cfg, err := config.LoadDefaultConfig(ctx)
	if err != nil {
		return nil, err
	}

	fmt.Println("region", cfg.Region)
	fmt.Println("endpoint", endpoint)

	mediaConvertClient := mediaconvert.NewFromConfig(cfg, func(o *mediaconvert.Options) {
		o.Region = cfg.Region
		o.BaseEndpoint = &endpoint
	})

	return mediaConvertClient, nil
}
