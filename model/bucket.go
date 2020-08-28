package model

import (
	"bytes"
	"couchbasecoreapi/config"
	view_model "couchbasecoreapi/view-model"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
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
func (s *BucketModel) CreateNewBucket(newBucket view_model.BucketEditCommand) (statusCode int, err error) {
	data := url.Values{}
	data.Set("name", newBucket.Name)
	data.Set("ramQuotaMB", newBucket.RamQuotaMB)
	data.Set("bucketType", newBucket.BucketType)
	data.Set("evictionPolicy", newBucket.EvictionPolicy)
	data.Set("durabilityMinLevel", newBucket.DurabilityMinLevel)
	data.Set("threadsNumber", newBucket.ThreadsNumber)
	data.Set("replicaNumber", newBucket.ReplicaNumber)
	data.Set("replicaIndex", newBucket.ReplicaIndex)
	data.Set("conflictResolutionType", newBucket.ConflictResolutionType)
	data.Set("flushEnabled", newBucket.FlushEnabled)
	data.Set("autoCompactionDefined", newBucket.AutoCompactionDefined)
	data.Set("parallelDBAndViewCompaction", newBucket.ParallelDBAndViewCompaction)
	data.Set("databaseFragmentationThreshold[percentage]", newBucket.DatabaseFragmentationThresholdPercentage)
	data.Set("databaseFragmentationThreshold[size]", newBucket.DatabaseFragmentationThresholdSize)
	data.Set("viewFragmentationThreshold[percentage]", newBucket.ViewFragmentationThresholdPercentage)
	data.Set("viewFragmentationThreshold[size]", newBucket.ViewFragmentationThresholdSize)
	data.Set("indexCompactionMode", newBucket.IndexCompactionMode)
	data.Set("purgeInterval", newBucket.PurgeInterval)
	data.Set("allowedTimePeriod[fromHour]", newBucket.AllowedTimePeriodFromHour)
	data.Set("allowedTimePeriod[fromMinute]", newBucket.AllowedTimePeriodFromMinute)
	data.Set("allowedTimePeriod[toHour]", newBucket.AllowedTimePeriodToHour)
	data.Set("allowedTimePeriod[toMinute]", newBucket.AllowedTimePeriodToMinute)
	data.Set("allowedTimePeriod[abortOutside]", newBucket.AllowedTimePeriodAbortOutside)
	var endPoint = config.GetConnectionString(newBucket.ConnectionString, newBucket.UserName, newBucket.Password) + "/pools/" + newBucket.ClusterName + "/buckets"
	req, err := http.NewRequest("POST", endPoint, bytes.NewBufferString(data.Encode()))
	if err != nil {
		log.Fatalf("Error Occured. %+v", err)
	}
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	response, err := config.HttpClient.Do(req)
	if err != nil && response == nil {
		log.Fatalf("Error sending request to API endpoint. %+v", err)
	} else {
		defer response.Body.Close()
		t, _ := ioutil.ReadAll(response.Body)
		var r interface{}
		//if err != nil {
		//	log.Fatalf("Couldn't parse response body. %+v", err)
		//}
		json.Unmarshal(t, &r)
		fmt.Print(t)
	}
	return statusCode, err
}
func (s *BucketModel) GetBucketByName(bucket view_model.BucketDetailSearchCommand) (result Bucket, err error) {
	connectionStr := config.GetConnectionString(bucket.ConnectionString, bucket.UserName, bucket.Password) + "/pools/" + bucket.ClusterName + "/buckets/" + bucket.BucketName
	response, err := http.Get(connectionStr)
	if err != nil {
		fmt.Print(err.Error())
		os.Exit(1)
	}
	responseData, err := ioutil.ReadAll(response.Body)
	er := json.Unmarshal(responseData, &result)
	return result, er
}
func (s *BucketModel) EditBucket(newBucket view_model.BucketEditCommand) (statusCode int, err error) {
	data := url.Values{}
	data.Set("ramQuotaMB", newBucket.RamQuotaMB)
	data.Set("bucketType", newBucket.BucketType)
	data.Set("evictionPolicy", newBucket.EvictionPolicy)
	data.Set("durabilityMinLevel", newBucket.DurabilityMinLevel)
	data.Set("threadsNumber", newBucket.ThreadsNumber)
	data.Set("replicaNumber", newBucket.ReplicaNumber)
	data.Set("replicaIndex", newBucket.ReplicaIndex)
	data.Set("conflictResolutionType", newBucket.ConflictResolutionType)
	data.Set("flushEnabled", newBucket.FlushEnabled)
	data.Set("autoCompactionDefined", newBucket.AutoCompactionDefined)
	data.Set("parallelDBAndViewCompaction", newBucket.ParallelDBAndViewCompaction)
	data.Set("databaseFragmentationThreshold[percentage]", newBucket.DatabaseFragmentationThresholdPercentage)
	data.Set("databaseFragmentationThreshold[size]", newBucket.DatabaseFragmentationThresholdSize)
	data.Set("viewFragmentationThreshold[percentage]", newBucket.ViewFragmentationThresholdPercentage)
	data.Set("viewFragmentationThreshold[size]", newBucket.ViewFragmentationThresholdSize)
	data.Set("indexCompactionMode", newBucket.IndexCompactionMode)
	data.Set("purgeInterval", newBucket.PurgeInterval)
	data.Set("allowedTimePeriod[fromHour]", newBucket.AllowedTimePeriodFromHour)
	data.Set("allowedTimePeriod[fromMinute]", newBucket.AllowedTimePeriodFromMinute)
	data.Set("allowedTimePeriod[toHour]", newBucket.AllowedTimePeriodToHour)
	data.Set("allowedTimePeriod[toMinute]", newBucket.AllowedTimePeriodToMinute)
	data.Set("allowedTimePeriod[abortOutside]", newBucket.AllowedTimePeriodAbortOutside)
	var endPoint = config.GetConnectionString(newBucket.ConnectionString, newBucket.UserName, newBucket.Password) + "/pools/" + newBucket.ClusterName + "/buckets/" + newBucket.Name
	req, err := http.NewRequest("PUT", endPoint, bytes.NewBufferString(data.Encode()))
	if err != nil {
		log.Fatalf("Error Occured. %+v", err)
	}
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	response, err := config.HttpClient.Do(req)
	if err != nil && response == nil {
		log.Fatalf("Error sending request to API endpoint. %+v", err)
	} else {
		defer response.Body.Close()
		t, _ := ioutil.ReadAll(response.Body)
		var r interface{}
		json.Unmarshal(t, &r)
		fmt.Print(t)
	}
	return statusCode, err
}
func (s *BucketModel) DeleteBucket(bucket view_model.BucketDetailSearchCommand) (statusCode int, err error) {
	connectionStr := config.GetConnectionString(bucket.ConnectionString, bucket.UserName, bucket.Password) + "/pools/" + bucket.ClusterName + "/buckets/" + bucket.BucketName
	response, err := http.Get(connectionStr)
	if err != nil {
		fmt.Print(err.Error())
		os.Exit(1)
	}
	statusCode = response.StatusCode
	return statusCode, err
}
