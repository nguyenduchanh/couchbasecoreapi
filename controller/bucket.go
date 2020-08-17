package controller

import (
	"couchbasecoreapi/model"
	view_model "couchbasecoreapi/view-model"
	"github.com/gin-gonic/gin"
)

var bucketModel = new(model.BucketModel)

type BucketController struct{}

func (b *BucketController) SelectAll(c *gin.Context) {
	var data view_model.BucketSearchCommand
	if c.BindJSON(&data) != nil {
		c.JSON(406, gin.H{"message": "Invalid form", "form": data})
		c.Abort()
		return
	}
	list, err := bucketModel.SearchBucket(data)
	if err != nil {
		c.JSON(404, gin.H{"message": "Get bucket error", "error": err.Error()})
		c.Abort()
	} else {
		c.JSON(200, gin.H{"data": list})
	}
}
