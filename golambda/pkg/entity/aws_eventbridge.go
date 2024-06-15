package entity

type AWSEventBridgeS3Event struct {
	Detail     AWSEventBridgeS3Detail `json:"detail"`
	DetailType string                 `json:"detail-type"`
	ID         string                 `json:"id"`
	Region     string                 `json:"region"`
	Resources  []string               `json:"resources"`
	Source     string                 `json:"source"`
}

type AWSEventBridgeS3Detail struct {
	Bucket AWSEventBridgeS3DetailBucket `json:"bucket"`
	Object AWSEventBridgeS3DetailObject `json:"object"`
	Reason string                       `json:"reason"`
}

type AWSEventBridgeS3DetailBucket struct {
	Name string `json:"name"`
}

type AWSEventBridgeS3DetailObject struct {
	Key  string `json:"key"`
	Size int64  `json:"size"`
}
