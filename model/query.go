package model

import (
	"bytes"
	"couchbasecoreapi/config"
	view_model "couchbasecoreapi/view-model"
	"encoding/json"
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
	Errors    interface{}   `json:"errors"`
}
type Metrics struct {
	ElapsedTime   string `json:"elapsedTime"`
	ExecutionTime string `json:"executionTime"`
	ResultCount   string `json:"resultCount"`
	ResultSize    string `json:"resultSize"`
}
type QueryModel struct{}

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
func (s *QueryModel) SelectAll(connectionString string, userName string, password string, query string) (_query interface{}, error error) {
	data := url.Values{}
	data.Set("statement", query)
	var endPoint = config.GetQueryConnectionString(connectionString, userName, password) + "/query/service"
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
	}
	return _query, nil
}
func (s *QueryModel) QueryExecute(query view_model.QueryExecuteCommand) (_query interface{}, error error) {
	data := url.Values{}
	data.Set("statement", query.Query)
	var endPoint = config.GetQueryConnectionString(query.ConnectionString, query.UserName, query.Password) + "/query/service"
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
	}
	return _query, nil
}
