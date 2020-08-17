package view_model

type BucketCreateCommand struct {
	Name            string `json:"name"` // new bucket name
	RamQuotaMB      string `json:"ramQuotaMB"`
	AuthType        string `json:"authType"` //none | sasl
	ReplicaNumber   string `json:"replicaNumber"`
	ProxyPort       string `json:"proxyPort"`
	BucketType      string `json:"bucketType"`
	CompressionMode string `json:"compressionMode"`
	MaxTTL          string `json:"maxTTL"`
	FlushEnabled    string `json:"flushEnabled"`
	ReplicaIndex    string `json:"replicaIndex"`
	ThreadsNumber   string `json:"threadsNumber"`
	EvictionPolicy  string `json:"evictionPolicy"`
}
type BucketSearchCommand struct {
	ConnectionString string `json:"connectionstring" binding:"required"`
	UserName         string `json:"username" binding:"required"`
	Password         string `json:"password" binding:"required"`
	ClusterName      string `json:"clustername" binding:"required"`
}
