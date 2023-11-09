# cf-sam-video-transcription-translate
AWS SAM (Serverless Application Model) that performs the followings:
1. Convert mp4 (video s3 bucket) to mp3 file (audio s3 bucket) using AWS MediaConvert
2. Transcribe mp3 file (audio s3 bucket) to subtitle (SRT) file (transcription s3 bucket) using AWS Transcribe
3. Translate subtitle (SRT) file (transcription s3 bucket) from the source to target language (translation s3 bucket) using AWS Translate

## Getting Started

## Local Deployment

### Run on local machine
```
npm i -g cfn-include
make build
sam deploy --guided
```
