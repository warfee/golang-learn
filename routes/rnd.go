package routes

import (
	"github.com/gin-gonic/gin"
	"proj1/handlers"
)

func RndRoutes(router *gin.Engine) {
	rndRouteGroup := router.Group("/rnd")
	{
		rndRouteGroup.GET("/ping", handlers.Ping)
		rndRouteGroup.GET("/stats", handlers.Stats)
		rndRouteGroup.GET("/queue-slack", handlers.QueueSlack)

		// rndRouteGroup.GET("/push-queue", handlers.PushQueue)
	}
}