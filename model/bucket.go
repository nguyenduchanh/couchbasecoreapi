package model

import (
	"couchbasecoreapi/config"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
)

type Bucket struct {
	Name                   string `json:"name"`
	Uuid                   string `json:"uuid"`
	BucketType             string `json:"bucketType"`
	AuthType               string `json:"authType"`
	Uri                    string `json:"uri"`
	StreamingUri           string `json:"streamingUri"`
	LocalRandomKeyUri      string `json:"localRandomKeyUri"`
	Controllers            string `json:"controllers"`
	Nodes                  string `json:"nodes"`
	Stats                  string `json:"stats"`
	NodeLocator            string `json:"nodeLocator"`
	SaslPassword           string `json:"saslPassword"`
	Ddocs                  string `json:"ddocs"`
	ReplicaIndex           bool   `json:"replicaIndex"`
	AutoCompactionSettings bool   `json:"autoCompactionSettings"`
	VBucketServerMap       string `json:"vBucketServerMap"`
	ReplicaNumber          int    `json:"replicaNumber"`
	ThreadsNumber          int    `json:"threadsNumber"`
	Quota                  string `json:"quota"`
	BasicStats             string `json:"basicStats"`
	EvictionPolicy         string `json:"evictionPolicy"`
	ConflictResolutionType string `json:"conflictResolutionType"`
	BucketCapabilitiesVer  string `json:"bucketCapabilitiesVer"`
	BucketCapabilities     string `json:"bucketCapabilities"`
}
type BucketModel struct{}

func (s *BucketModel) GetAllBucket(clusterName string) (data []Bucket, error error) {
	connectionString := config.GetConnectionString() + "/pools/" + clusterName + "/buckets"
	response, err := http.Get(connectionString)
	if err != nil {
		fmt.Print(err.Error())
		os.Exit(1)
	}
	responseData, err := ioutil.ReadAll(response.Body)
	//var b Bucket
	json.Unmarshal(responseData, &data)
	//var result []string
	//
	//for  i := int32(0); i < int32(len(data)); i++{
	//	result = append(result, data[i].Name)
	//}
	return data, nil
}
