package eventbridge

type S3 struct {
	Detail     S3Detail `json:"detail"`
	DetailType string   `json:"detail-type"`
	ID         string   `json:"id"`
	Region     string   `json:"region"`
	Resources  []string `json:"resources"`
	Source     string   `json:"source"`
}

type S3Detail struct {
	Bucket S3DetailBucket `json:"bucket"`
	Object S3DetailObject `json:"object"`
	Reason string         `json:"reason"`
}

type S3DetailBucket struct {
	Name string `json:"name"`
}

type S3DetailObject struct {
	Key  string `json:"key"`
	Size int64  `json:"size"`
}
