package controller

import (
	"couchbasecoreapi/model"
	view_model "couchbasecoreapi/view-model"
	"github.com/gin-gonic/gin"
)

var clusterModel = new(model.ClusterModel)

type ClusterController struct{}

func (b *ClusterController) SelectAllCluster(c *gin.Context) {
	var data view_model.ClusterSearchCommand
	if c.BindJSON(&data) != nil {
		c.JSON(406, gin.H{"message": "Invalid form", "form": data})
		c.Abort()
		return
	}
	list, err := clusterModel.SelectAllCluster(data)
	if err != nil {
		c.JSON(404, gin.H{"message": "Get cluster error", "error": err.Error()})
		c.Abort()
	} else {
		c.JSON(200, gin.H{"data": list})
	}
}
