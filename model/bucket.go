package model

import (
	"couchbasecoreapi/config"
	view_model "couchbasecoreapi/view-model"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"os"
)

type Bucket struct {
	Name                   string      `json:"name"`
	Uuid                   string      `json:"uuid"`
	RamQuotaMB             int         `json:"ramQuotaMB"`
	BucketType             string      `json:"bucketType"`
	AuthType               string      `json:"authType"`
	Uri                    string      `json:"uri"`
	StreamingUri           string      `json:"streamingUri"`
	LocalRandomKeyUri      string      `json:"localRandomKeyUri"`
	Controllers            interface{} `json:"controllers"`
	Nodes                  interface{} `json:"nodes"`
	Stats                  interface{} `json:"stats"`
	NodeLocator            string      `json:"nodeLocator"`
	SaslPassword           string      `json:"saslPassword"`
	Ddocs                  interface{} `json:"ddocs"`
	ReplicaIndex           bool        `json:"replicaIndex"`
	AutoCompactionSettings bool        `json:"autoCompactionSettings"`
	VBucketServerMap       interface{} `json:"vBucketServerMap"`
	ReplicaNumber          int         `json:"replicaNumber"`
	ThreadsNumber          int         `json:"threadsNumber"`
	Quota                  interface{} `json:"quota"`
	BasicStats             interface{} `json:"basicStats"`
	EvictionPolicy         string      `json:"evictionPolicy"`
	ConflictResolutionType string      `json:"conflictResolutionType"`
	BucketCapabilitiesVer  string      `json:"bucketCapabilitiesVer"`
	BucketCapabilities     interface{} `json:"bucketCapabilities"`
}
type BucketModel struct{}

func (s *BucketModel) SearchBucket(bucket view_model.BucketSearchCommand) (data []Bucket, error error) {
	connectionStr := config.GetConnectionString(bucket.ConnectionString, bucket.UserName, bucket.Password) + "/pools/" + bucket.ClusterName + "/buckets"
	response, err := http.Get(connectionStr)
	if err != nil {
		fmt.Print(err.Error())
		os.Exit(1)
	}
	responseData, err := ioutil.ReadAll(response.Body)
	er := json.Unmarshal(responseData, &data)
	return data, er
}
func (s *BucketModel) CreateNewBucket(newBucket view_model.BucketCreateCommand) (statusCode int, err error) {
	//connectionStr := config.GetConnectionString(newBucket.ConnectionString, newBucket.UserName, newBucket.Password) + "/pools/" + newBucket.ClusterName + "/buckets"
	data := url.Values{}
	data.Set("name", newBucket.Name)
	data.Set("ramQuotaMB", string(newBucket.RamQuotaMB))
	data.Set("bucketType", newBucket.BucketType)
	data.Set("evictionPolicy", newBucket.EvictionPolicy)
	data.Set("durabilityMinLevel", newBucket.DurabilityMinLevel)
	data.Set("threadsNumber", newBucket.ThreadsNumber)
	data.Set("replicaNumber", newBucket.ReplicaNumber)
	data.Set("threadsNumber", newBucket.ThreadsNumber)
	data.Set("compressionMode", newBucket.CompressionMode)
	data.Set("maxTTL", newBucket.MaxTTL)
	data.Set("replicaIndex", newBucket.ReplicaIndex)
	data.Set("conflictResolutionType", newBucket.ConflictResolutionType)
	data.Set("flushEnabled", newBucket.FlushEnabled)
	data.Set("autoCompactionDefined", newBucket.AutoCompactionDefined)
	data.Set("parallelDBAndViewCompaction", newBucket.AutoCompactionDefined)

	return statusCode, err
}
