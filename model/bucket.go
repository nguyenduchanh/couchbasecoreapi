package model

import (
	"couchbasecoreapi/config"
	view_model "couchbasecoreapi/view-model"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
)

type Bucket struct {
	Name                   string      `json:"name"`
	Uuid                   string      `json:"uuid"`
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
func (s *BucketModel) CreateNewBucket(clusterName string) (statusCode int, err error) {
	return 202, nil
}
