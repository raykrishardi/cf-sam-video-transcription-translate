package transcribe

import (
	"context"
	"fmt"

	"github.com/raykrishardi/cf-sam-video-transcription-translate/golambda/pkg/utils"

	"github.com/raykrishardi/cf-sam-video-transcription-translate/golambda/pkg/entity"

	"github.com/raykrishardi/cf-sam-video-transcription-translate/golambda/config"

	"github.com/aws/aws-sdk-go-v2/service/transcribe"
	"github.com/aws/aws-sdk-go-v2/service/transcribe/types"
)

type AWSTranscribeRepo struct {
	App    *config.AppConfig
	Client *transcribe.Client
}

func NewAWSTranscribeRepo(app *config.AppConfig, client *transcribe.Client) *AWSTranscribeRepo {
	return &AWSTranscribeRepo{
		App:    app,
		Client: client,
	}
}

func (r *AWSTranscribeRepo) TranscribeMP3ToSRT(ctx context.Context, params entity.TranscribeMP3ToSRTInput) error {
	// Specify the starting value that is assigned to the first subtitle segment. The
	// default start index for Amazon Transcribe is 0 , which differs from the more
	// widely used standard of 1 . If you're uncertain which value to use, we recommend
	// choosing 1 , as this may improve compatibility with other services.
	defaultSubtitleOutputStartIndex := int32(1)

	jobName := fmt.Sprintf("%s-%s", params.InFileName, utils.GetRandomString(5))
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

	_, err := r.Client.StartTranscriptionJob(ctx, stji)
	if err != nil {
		return err
	}

	return nil
}
