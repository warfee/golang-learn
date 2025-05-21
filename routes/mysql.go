package routes

import (
	"github.com/gin-gonic/gin"
	"proj1/handlers"
)

func MysqlRoutes(router *gin.Engine) {
	mysqlRouteGroup := router.Group("/rnd/mysql")
	{
		mysqlRouteGroup.GET("/one", handlers.MysqlSelectOne)
		mysqlRouteGroup.GET("/insert", handlers.MysqlInsert)
		mysqlRouteGroup.GET("/update", handlers.MysqlUpdate)
		mysqlRouteGroup.GET("/delete", handlers.MysqlDelete)
	}
}