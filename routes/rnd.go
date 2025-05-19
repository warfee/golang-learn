package routes

import (
	"github.com/gin-gonic/gin"
	"proj1/controllers"
)

func RndRoutes(router *gin.Engine) {
	rndRouteGroup := router.Group("/rnd")
	{
		rndRouteGroup.GET("/ping", controllers.Ping)
		rndRouteGroup.GET("/stats", controllers.Stats)
	}
}