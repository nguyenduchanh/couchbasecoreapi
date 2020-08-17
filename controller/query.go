package controller

import (
	"couchbasecoreapi/model"
	view_model "couchbasecoreapi/view-model"
	"github.com/gin-gonic/gin"
)

var queryModel = new(model.QueryModel)

type QueryController struct{}

func (q *QueryController) SelectAll(c *gin.Context) {
	connectionString := c.Param("connectionstring")
	userName := c.Param("username")
	password := c.Param("password")
	query := c.Param("query")

	list, err := queryModel.SelectAll(connectionString, userName, password, query)
	if err != nil {
		c.JSON(404, gin.H{"message": "Find Error", "error": err.Error()})
		c.Abort()
	} else {
		c.JSON(200, gin.H{"data": list})
	}
}
func (q *QueryController) QueryExecute(c *gin.Context) {
	var data view_model.QueryExecuteCommand
	if c.BindJSON(&data) != nil {
		c.JSON(406, gin.H{"message": "Invalid form", "form": data})
		c.Abort()
		return
	}
	list, err := queryModel.QueryExecute(data)
	if err != nil {
		c.JSON(404, gin.H{"message": "Execute query Error", "error": err.Error()})
		c.Abort()
	} else {
		c.JSON(200, gin.H{"data": list})
	}
}
