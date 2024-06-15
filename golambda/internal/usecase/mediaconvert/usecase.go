package mediaconvert

import (
	mcrepo "cf-sam-video-transcription-translate/internal/pkg/mediaconvert"
	"context"

	"github.com/aws/aws-sdk-go-v2/service/mediaconvert"
	"github.com/aws/aws-sdk-go-v2/service/mediaconvert/types"
)

type MediaConvertUseCase struct {
	MediaConvertRepo *mcrepo.MediaConvertRepository
	Client           *mediaconvert.Client
}

func NewMediaConvertUseCase(ctx context.Context, mcRepo *mcrepo.MediaConvertRepository) *MediaConvertUseCase {
	uc := &MediaConvertUseCase{}

	client, err := mcRepo.GetMediaConvertClient(ctx)
	if err != nil {
		return uc
	}

	uc.MediaConvertRepo = mcRepo
	uc.Client = client

	return uc
}

func (uc *MediaConvertUseCase) ConvertMP4ToMP3(ctx context.Context, params ConvertMP4ToMP3Input) (*mediaconvert.CreateJobOutput, error) {
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

	return uc.Client.CreateJob(ctx, cji)
}
