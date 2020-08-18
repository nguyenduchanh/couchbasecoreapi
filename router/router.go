package router

import (
	"couchbasecoreapi/controller"
	"couchbasecoreapi/model"
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"net/http"
	"os"
)

var portModel = new(model.PortModel)

func CORS() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT, DELETE")
		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}

		c.Next()
	}
}
func Router() {
	port, _ := portModel.ReadPortJson()
	fmt.Print(port)
	router := gin.Default()
	router.Use(CORS())
	v1 := router.Group("/v1")
	{
		//add controller
		bucket := new(controller.BucketController)
		cluster := new(controller.ClusterController)
		query := new(controller.QueryController)
		//add bucket router
		v1.POST("/buckets", bucket.SelectAll)
		//add cluster router
		v1.POST("/clusters", cluster.SelectAllCluster)
		//add query router
		v1.POST("/query", query.QueryExecute)
	}
	router.NoRoute(func(c *gin.Context) {
		c.String(http.StatusNotFound, "Not Found")
	})
	if port.Port != "" {
		router.Run(":" + port.Port)
	} else {
		p := new(model.PortJson)
		p.Port = "8999"
		jsonString, _ := json.Marshal(p)
		ioutil.WriteFile("port.json", jsonString, os.ModePerm)
		router.Run(":8999")
	}

}
