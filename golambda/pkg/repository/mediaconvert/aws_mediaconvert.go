package mediaconvert

import (
	"context"

	"github.com/raykrishardi/cf-sam-video-transcription-translate/golambda/pkg/entity"

	"github.com/raykrishardi/cf-sam-video-transcription-translate/golambda/config"

	"github.com/aws/aws-sdk-go-v2/service/mediaconvert"
	"github.com/aws/aws-sdk-go-v2/service/mediaconvert/types"
)

type AWSMediaConvertRepo struct {
	App    *config.AppConfig
	Client *mediaconvert.Client
}

func NewAWSMediaConvertRepo(app *config.AppConfig, client *mediaconvert.Client) *AWSMediaConvertRepo {
	return &AWSMediaConvertRepo{
		App:    app,
		Client: client,
	}
}

func (r *AWSMediaConvertRepo) ConvertMP4ToMP3(ctx context.Context, params entity.ConvertMP4ToMP3Input) error {
	inAudioSelectors := map[string]types.AudioSelector{
		"Audio Selector 1": {
			DefaultSelection: types.AudioDefaultSelectionDefault,
		},
	}
	cji := &mediaconvert.CreateJobInput{
		Role: &params.Role,
		Settings: &types.JobSettings{
			Inputs: []types.Input{
				{
					FileInput:      &params.InS3Uri,
					AudioSelectors: inAudioSelectors,
				},
			},
			OutputGroups: []types.OutputGroup{
				{
					Outputs: []types.Output{
						{
							ContainerSettings: &types.ContainerSettings{
								Container: types.ContainerTypeRaw,
							},
							AudioDescriptions: []types.AudioDescription{
								{
									CodecSettings: &types.AudioCodecSettings{
										Codec: types.AudioCodecMp3,
										Mp3Settings: &types.Mp3Settings{
											RateControlMode: types.Mp3RateControlMode(params.OutMP3Settings.RateControlMode),
											Bitrate:         params.OutMP3Settings.BitRate,
											VbrQuality:      params.OutMP3Settings.VbrQuality,
										},
									},
								},
							},
						},
					},
					OutputGroupSettings: &types.OutputGroupSettings{
						Type: types.OutputGroupTypeFileGroupSettings,
						FileGroupSettings: &types.FileGroupSettings{
							Destination: &params.OutS3Uri,
						},
					},
				},
			},
		},
	}

	_, err := r.Client.CreateJob(ctx, cji)
	if err != nil {
		return err
	}

	return nil
}
