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

type Cluster struct {
	IsAdminCreds          bool   `json:"isAdminCreds"`
	IsROAdminCreds        bool   `json:"isROAdminCreds"`
	IsEnterprise          bool   `json:"isEnterprise"`
	AllowedServices       string `json:"allowedServices"`
	IsIPv6                bool   `json:"isIPv6"`
	IsDeveloperPreview    bool   `json:"isDeveloperPreview"`
	Pools                 []Pool `json:"pools"`
	Settings              string `json:"settings"`
	Uuid                  string `json:"uuid"`
	ImplementationVersion string `json:"implementationVersion"`
	ComponentsVersion     string `json:"componentsVersion"`
}
type Pool struct {
	Name         string `json:"name"`
	Uri          string `json:"uri"`
	StreamingUri string `json:"streamingUri"`
}
type ClusterModel struct {
}

func (c *ClusterModel) SelectAllCluster(cluster view_model.ClusterSearchCommand) (data Cluster, error error) {
	connectionStr := config.GetConnectionString(cluster.ConnectionString, cluster.UserName, cluster.Password)
	response, err := http.Get(connectionStr + "/pools")
	if err != nil {
		fmt.Print(err.Error())
		os.Exit(1)
	}
	responseData, err := ioutil.ReadAll(response.Body)
	json.Unmarshal(responseData, &data)
	fmt.Print(data)
	return data, nil
}
