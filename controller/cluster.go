package controller

import (
	"couchbasecoreapi/model"
	"github.com/gin-gonic/gin"
)

var clusterModel = new(model.ClusterModel)

type ClusterController struct{}

func (cl *ClusterController) GetCluster(c *gin.Context) {
	list, err := clusterModel.GetAllCluster()
	if err != nil {
		c.JSON(404, gin.H{"message": "Find Error", "error": err.Error()})
		c.Abort()
	} else {
		c.JSON(200, gin.H{"data": list})
	}
}
