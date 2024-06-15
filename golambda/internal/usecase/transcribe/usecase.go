package transcribe

import (
	"cf-sam-video-transcription-translate/internal/pkg/helper"
	trrepo "cf-sam-video-transcription-translate/internal/pkg/transcribe"
	"context"
	"fmt"

	"github.com/aws/aws-sdk-go-v2/service/transcribe"
	"github.com/aws/aws-sdk-go-v2/service/transcribe/types"
)

type TranscribeUseCase struct {
	TranscribeRepo *trrepo.TranscribeRepository
	Client         *transcribe.Client
}

func NewTranscribeUseCase(ctx context.Context, trRepo *trrepo.TranscribeRepository) *TranscribeUseCase {
	uc := &TranscribeUseCase{}

	client, err := trRepo.GetTranscribeClient(ctx)
	if err != nil {
		return uc
	}

	uc.TranscribeRepo = trRepo
	uc.Client = client

	return uc
}

func (uc *TranscribeUseCase) TranscribeMP3ToSRT(ctx context.Context, params TranscribeMP3ToSRTInput) (*transcribe.StartTranscriptionJobOutput, error) {
	// Specify the starting value that is assigned to the first subtitle segment. The
	// default start index for Amazon Transcribe is 0 , which differs from the more
	// widely used standard of 1 . If you're uncertain which value to use, we recommend
	// choosing 1 , as this may improve compatibility with other services.
	defaultSubtitleOutputStartIndex := int32(1)

	jobName := fmt.Sprintf("%s-%s", params.InFileName, helper.RandomString(5))
	outputKeyJSON := fmt.Sprintf("%s.json", *params.OutBucketObjectKey)
	stji := &transcribe.StartTranscriptionJobInput{
		Media: &types.Media{
			MediaFileUri: &params.InS3Uri,
		},
		MediaFormat:          types.MediaFormatMp3,
		TranscriptionJobName: &jobName,
		IdentifyLanguage:     params.IdentifyLanguage,
		OutputBucketName:     &params.OutBucketName,
		OutputKey:            &outputKeyJSON,
		Subtitles: &types.Subtitles{
			Formats:          []types.SubtitleFormat{types.SubtitleFormatSrt},
			OutputStartIndex: &defaultSubtitleOutputStartIndex,
		},
	}

	return uc.Client.StartTranscriptionJob(ctx, stji)
}
