package s3

import (
	s3repo "cf-sam-video-transcription-translate/internal/pkg/s3"
	"context"

	"github.com/aws/aws-sdk-go-v2/service/s3"
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

func (uc *S3UseCase) GetObject(ctx context.Context, params GetObjectInput) (*s3.GetObjectOutput, error) {
	goi := &s3.GetObjectInput{
		Bucket: &params.BucketName,
		Key:    &params.Key,
	}

	return uc.Client.GetObject(ctx, goi)
}
