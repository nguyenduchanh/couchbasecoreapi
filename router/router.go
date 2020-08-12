package router

import (
	"couchbasecoreapi/controller"
	"github.com/gin-gonic/gin"
	"net/http"
)

func Router() {
	router := gin.Default()
	v1 := router.Group("/v1")
	{
		//add controller
		bucket := new(controller.BucketController)
		cluster := new(controller.ClusterController)
		//add bucket router
		v1.GET("/buckets/:clustername", bucket.GetAllBucket)
		//add cluster router
		v1.GET("/clusters", cluster.GetCluster)
	}
	router.NoRoute(func(c *gin.Context) {
		c.String(http.StatusNotFound, "Not Found")
	})

	router.Run(":8080")
}
