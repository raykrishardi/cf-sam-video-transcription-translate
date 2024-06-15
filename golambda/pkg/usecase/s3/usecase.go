package s3

import (
	"bytes"
	s3repo "cf-sam-video-transcription-translate/pkg/repository/s3"
	"context"

	"github.com/aws/aws-sdk-go-v2/service/s3"

	"cf-sam-video-transcription-translate/pkg/entity"
)

type S3UseCase struct {
	S3Repo *s3repo.S3Repository
	Client *s3.Client
}

func NewS3UseCase(ctx context.Context, s3Repo *s3repo.S3Repository) *S3UseCase {
	uc := &S3UseCase{}

	client, err := s3Repo.GetS3Client(ctx)
	if err != nil {
		return uc
	}

	uc.S3Repo = s3Repo
	uc.Client = client

	return uc
}

func (uc *S3UseCase) GetObject(ctx context.Context, params entity.GetObjectInput) (*s3.GetObjectOutput, error) {
	goi := &s3.GetObjectInput{
		Bucket: &params.BucketName,
		Key:    &params.Key,
	}

	return uc.Client.GetObject(ctx, goi)
}

func (uc *S3UseCase) PutObject(ctx context.Context, params entity.PutObjectInput) (*s3.PutObjectOutput, error) {
	poi := &s3.PutObjectInput{
		Bucket: &params.BucketName,
		Key:    &params.Key,
		Body:   bytes.NewReader(params.Body),
	}

	return uc.Client.PutObject(ctx, poi)
}
