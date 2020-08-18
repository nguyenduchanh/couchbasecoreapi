package config

import (
	"net/http"
	"time"
)

var (
	HttpClient *http.Client
)

const (
	MaxIdleConnections int = 20
	RequestTimeout     int = 5
)

// init HTTPClient
func init() {
	HttpClient = createHTTPClient()
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
