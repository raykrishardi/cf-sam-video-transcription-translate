package objectstore

import (
	"bytes"
	"cf-sam-video-transcription-translate/config"
	"cf-sam-video-transcription-translate/pkg/entity"
	"context"
	"io"

	"github.com/aws/aws-sdk-go-v2/service/s3"
)

type S3Repo struct {
	App    *config.AppConfig
	Client *s3.Client
}

func NewS3Repo(app *config.AppConfig, client *s3.Client) *S3Repo {
	return &S3Repo{
		App:    app,
		Client: client,
	}
}

func (r *S3Repo) GetObject(ctx context.Context, params entity.GetObjectInput) ([]byte, error) {
	goi := &s3.GetObjectInput{
		Bucket: &params.BucketName,
		Key:    &params.Key,
	}

	goo, err := r.Client.GetObject(ctx, goi)
	if err != nil {
		return nil, err
	}

	content, err := io.ReadAll(goo.Body)
	if err != nil {
		return nil, err
	}

	return content, nil
}

func (r *S3Repo) PutObject(ctx context.Context, params entity.PutObjectInput) error {
	poi := &s3.PutObjectInput{
		Bucket: &params.BucketName,
		Key:    &params.Key,
		Body:   bytes.NewReader(params.Body),
	}

	_, err := r.Client.PutObject(ctx, poi)
	if err != nil {
		return err
	}

	return nil
}
