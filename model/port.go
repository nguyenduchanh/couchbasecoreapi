package model

import (
	"encoding/json"
	"io/ioutil"
	"os"
)

type PortJson struct {
	Port string `json:"port"`
}
type PortModel struct{}

func (c *PortModel) ReadPortJson() (port *PortJson, err error) {
	if fileExists("port.json") {
		jsonFile, err := os.Open("port.json")
		if err != nil {
			p := new(PortJson)
			p.Port = "8999"
			return p, err
		}
		defer jsonFile.Close()
		byteValue, _ := ioutil.ReadAll(jsonFile)
		er := json.Unmarshal(byteValue, &port)
		return port, er
	}
	p := new(PortJson)
	p.Port = "8999"
	return p, err
}

// try using it to prevent further errors.
func fileExists(filename string) bool {
	info, err := os.Stat(filename)
	if os.IsNotExist(err) {
		return false
	}
	return !info.IsDir()
}
