package controller

import (
	"couchbasecoreapi/model"
	"github.com/gin-gonic/gin"
)

var clusterModel = new(model.ClusterModel)

type ClusterController struct{}

func (cl *ClusterController) GetCluster(c *gin.Context) {
	connectionString := c.Param("connectionstring")
	userName := c.Param("username")
	password := c.Param("password")
	list, err := clusterModel.GetAllCluster(connectionString, userName, password)
	if err != nil {
		c.JSON(404, gin.H{"message": "Find Error", "error": err.Error()})
		c.Abort()
	} else {
		c.JSON(200, gin.H{"data": list})
	}
}
