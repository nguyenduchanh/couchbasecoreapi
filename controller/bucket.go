package controller

import (
	"couchbasecoreapi/model"
	"github.com/gin-gonic/gin"
)

var bucketModel = new(model.BucketModel)

type BucketController struct{}

func (b *BucketController) GetAllBucket(c *gin.Context) {
	connectionString := c.Param("connectionstring")
	userName := c.Param("username")
	password := c.Param("password")
	clusterName := c.Param("clustername")
	list, err := bucketModel.GetAllBucket(connectionString, userName, password, clusterName)
	if err != nil {
		c.JSON(404, gin.H{"message": "Find Error", "error": err.Error()})
		c.Abort()
	} else {
		c.JSON(200, gin.H{"data": list})
	}
}
