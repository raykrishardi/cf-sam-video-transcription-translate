package s3

import (
	s3repo "cf-sam-video-transcription-translate/internal/pkg/s3"
	"context"

	"github.com/aws/aws-sdk-go-v2/service/s3"
)

type S3UseCase struct {
	S3Repo *s3repo.S3Repository
}

func NewS3UseCase(s3Repo *s3repo.S3Repository) *S3UseCase {
	return &S3UseCase{
		S3Repo: s3Repo,
	}
}

func (uc *S3UseCase) ListBucket(ctx context.Context, name string) (*s3.ListObjectsV2Output, error) {
	client, err := uc.S3Repo.GetS3Client(ctx)
	if err != nil {
		return nil, err
	}

	result, err := client.ListObjectsV2(ctx, &s3.ListObjectsV2Input{
		Bucket: &name,
	})
	if err != nil {
		return nil, err
	}

	return result, nil
}
