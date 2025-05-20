package routes

import (
	"github.com/gin-gonic/gin"
	"proj1/handlers"
)

func MongoRoutes(router *gin.Engine) {
	mongoRouteGroup := router.Group("/rnd/mongo")
	{
		mongoRouteGroup.GET("/one", handlers.MongoSelectOne)
		mongoRouteGroup.GET("/insert", handlers.MongoInsert)
		mongoRouteGroup.GET("/update", handlers.MongoUpdate)
		mongoRouteGroup.GET("/delete", handlers.MongoDelete)
	}
}