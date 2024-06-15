# cf-sam-video-transcription-translate
AWS SAM (Serverless Application Model) that performs the followings:
1. Convert mp4 (video s3 bucket) to mp3 file (audio s3 bucket) using AWS MediaConvert
2. Transcribe mp3 file (audio s3 bucket) to subtitle (SRT) file (transcription s3 bucket) using AWS Transcribe
3. Translate subtitle (SRT) file (transcription s3 bucket) from the source to target language (translation s3 bucket) using AWS Translate
    - If language source code for translation is set to "auto", AWS Comprehend will be used

## Getting Started

## Local Deployment

### Run on local machine
```
npm i -g cfn-include
make build

# Deploy the same app by specifying the vars
# This will create the specified bucket names
# Example value for MediaConvertEndpoint=https://mediaconvert.ap-southeast-2.amazonaws.com (reference from https://docs.aws.amazon.com/general/latest/gr/mediaconvert.html)
sam deploy --guided
```
