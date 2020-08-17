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
		query := new(controller.QueryController)
		//add bucket router
		v1.GET("/buckets/:connectionstring/:username/:password/:clustername", bucket.GetAllBucket)
		//add cluster router
		v1.GET("/clusters/:connectionstring/:username/:password", cluster.GetCluster)
		//add query router
		v1.POST("/query/:connectionstring/:username/:password/:query", query.SelectAll)
	}
	router.NoRoute(func(c *gin.Context) {
		c.String(http.StatusNotFound, "Not Found")
	})

	router.Run(":8999")
}
