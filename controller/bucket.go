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
func (b *BucketController) CreateNewBucket(c *gin.Context) {
	var data view_model.BucketEditCommand
	if c.BindJSON(&data) != nil {
		c.JSON(406, gin.H{"message": "Invalid form", "form": data})
		c.Abort()
		return
	}
	list, err := bucketModel.CreateNewBucket(data)
	if err != nil {
		c.JSON(404, gin.H{"message": "Create bucket error", "error": err.Error()})
		c.Abort()
	} else {
		c.JSON(200, gin.H{"data": list})
	}
}
func (b *BucketController) GetBucketByName(c *gin.Context) {
	var data view_model.BucketDetailSearchCommand
	if c.BindJSON(&data) != nil {
		c.JSON(406, gin.H{"message": "Invalid bucket detail search  form", "form": data})
		c.Abort()
		return
	}
	list, err := bucketModel.GetBucketByName(data)
	if err != nil {
		c.JSON(404, gin.H{"message": "Get bucket error", "error": err.Error()})
		c.Abort()
	} else {
		c.JSON(200, gin.H{"data": list})
	}
}
func (b *BucketController) EditBucket(c *gin.Context) {
	var data view_model.BucketEditCommand
	if c.BindJSON(&data) != nil {
		c.JSON(406, gin.H{"message": "Invalid form", "form": data})
		c.Abort()
		return
	}
	list, err := bucketModel.EditBucket(data)
	if err != nil {
		c.JSON(404, gin.H{"message": "Edit bucket error", "error": err.Error()})
		c.Abort()
	} else {
		c.JSON(200, gin.H{"data": list})
	}
}

func (b *BucketController) DeleteBucket(c *gin.Context) {
	var data view_model.BucketDetailSearchCommand
	if c.BindJSON(&data) != nil {
		c.JSON(406, gin.H{"message": "Invalid form", "form": data})
		c.Abort()
		return
	}
	list, err := bucketModel.DeleteBucket(data)
	if err != nil {
		c.JSON(404, gin.H{"message": "Delete bucket error", "error": err.Error()})
		c.Abort()
	} else {
		c.JSON(200, gin.H{"data": list})
	}
}
