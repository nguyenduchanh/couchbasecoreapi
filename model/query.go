package model

import (
	"bytes"
	"encoding/json"
	gocb "gopkg.in/couchbase/gocb.v2"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"time"
)

type Query struct {
	RequestID string        `json:"requestID"`
	Signature interface{}   `json:"signature"`
	Results   []interface{} `json:"results"`
	Status    string        `json:"status"`
	Metrics   interface{}   `json:"metrics"`
}
type Metrics struct {
	ElapsedTime   string `json:"elapsedTime"`
	ExecutionTime string `json:"executionTime"`
	ResultCount   string `json:"resultCount"`
	ResultSize    string `json:"resultSize"`
}
type QueryModel struct{}

var bucket *gocb.Bucket
var (
	httpClient *http.Client
)

const (
	MaxIdleConnections int = 20
	RequestTimeout     int = 5
)

// init HTTPClient
func init() {
	httpClient = createHTTPClient()
}

// createHTTPClient for connection re-use
func createHTTPClient() *http.Client {
	client := &http.Client{
		Transport: &http.Transport{
			MaxIdleConnsPerHost: MaxIdleConnections,
		},
		Timeout: time.Duration(RequestTimeout) * time.Second,
	}

	return client
}

func (s *QueryModel) SelectAll(query string) (_query Query, error error) {
	//resp, err := http.PostForm("http://Administrator:abc123@localhost:8093/query/service",
	//	url.Values{"statement": {"SELECT * FROM `classroom`"}})
	//if err != nil {
	//	panic(err)
	//}
	//var res map[string]Query
	//
	//json.NewDecoder(resp.Body).Decode(&res)
	//
	//responseData, err := ioutil.ReadAll(resp.Body)
	//json.Unmarshal(responseData, &_query)
	data := url.Values{}
	data.Set("statement", "SELECT * FROM `classroom`")
	var endPoint string = "http://Administrator:abc123@localhost:8093/query/service"

	req, err := http.NewRequest("POST", endPoint, bytes.NewBufferString(data.Encode()))
	if err != nil {
		log.Fatalf("Error Occured. %+v", err)
	}
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	response, err := httpClient.Do(req)
	if err != nil && response == nil {
		log.Fatalf("Error sending request to API endpoint. %+v", err)
	} else {
		defer response.Body.Close()
		body, err := ioutil.ReadAll(response.Body)
		if err != nil {
			log.Fatalf("Couldn't parse response body. %+v", err)
		}
		json.Unmarshal(body, &_query)
		//log.Println("Response Body:", string(body))
	}
	return _query, nil

}
