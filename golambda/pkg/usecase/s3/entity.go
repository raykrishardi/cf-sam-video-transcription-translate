package s3

type GetObjectInput struct {
	BucketName string
	Key        string
}

type PutObjectInput struct {
	BucketName string
	Key        string
	Body       []byte
}
