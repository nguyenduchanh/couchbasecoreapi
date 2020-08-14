package config

import (
	"encoding/json"
	"fmt"
	"github.com/sirupsen/logrus"
	"io/ioutil"
	"os"
	"time"
)

var ConnectionString = "localhost"
var Port = "8091"
var QueryPort = "8093"
var UserName = "Administrator"
var Password = "abc123"

type Configuration struct {
	ConnectionString string   `json:"connection_string"`
	Port             string   `json:"port"`
	UserName         string   `json:"user_name"`
	Password         string   `json:"password"`
	BucketList       []string `json:"bucket_list"`
	Os               string   `json:"os"`
	LastUpdate       string   `json:"last_update"`
}

func GetConnectionString() string {
	//http://Administrator:abc123@localhost:8091
	connectString := "http://" + UserName + ":" + Password + "@" + ConnectionString + ":" + Port
	return connectString
}
func GetQueryConnectionString() string {
	connecString := "http://" + UserName + ":" + Password + "@" + ConnectionString + ":" + QueryPort
	return connecString
}
func WriteFile() {
	filename := "config.json"
	err := checkFile(filename)
	if err != nil {
		logrus.Error(err)
	}

	file, err := ioutil.ReadFile(filename)
	if err != nil {
		logrus.Error(err)
	}

	data := []Configuration{}

	// Here the magic happens!
	json.Unmarshal(file, &data)

	newStruct := &Configuration{
		ConnectionString: "localhost",
		Port:             "8091",
		UserName:         "Adminstrator",
		Password:         "abc123",
		Os:               "Window",
		LastUpdate:       time.Now().String(),
	}

	data = append(data, *newStruct)

	// Preparing the data to be marshalled and written.
	dataBytes, err := json.Marshal(data)
	if err != nil {
		logrus.Error(err)
	}

	err = ioutil.WriteFile(filename, dataBytes, 0644)
	if err != nil {
		logrus.Error(err)
	}
}
func ReadFile() {
	//cData, err := ioutil.ReadFile("config.json")
	file, err := ioutil.ReadFile("config.json")
	if err != nil {
		fmt.Fprintf(os.Stderr, "Failed read file: %s\n", err)
		os.Exit(1)
	}

	var s Configuration
	s.unmarshalJSON(file)
	fmt.Print(s)
}

func (o *Configuration) unmarshalJSON(data []byte) error {
	var v [6]string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	fmt.Print(v)
	o.ConnectionString = v[0]
	o.Port = v[1]
	o.UserName = v[2]
	o.Password = v[3]
	o.LastUpdate = v[4]
	o.Os = v[5]
	return nil
}
func checkFile(filename string) error {
	_, err := os.Stat(filename)
	if os.IsNotExist(err) {
		_, err := os.Create(filename)
		if err != nil {
			return err
		}
	}
	return nil
}
