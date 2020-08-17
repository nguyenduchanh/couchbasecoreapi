package controller

import (
	"couchbasecoreapi/model"
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
