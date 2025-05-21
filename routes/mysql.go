package routes

import (
	"github.com/gin-gonic/gin"
	"proj1/handlers"
)

func MysqlRoutes(router *gin.Engine) {
	mysqlRouteGroup := router.Group("/mysql")
	{
		mysqlRouteGroup.GET("/one", handlers.MysqlSelectOne)
		mysqlRouteGroup.GET("/insert", handlers.MysqlInsert)
		mysqlRouteGroup.GET("/update", handlers.MysqlUpdate)
		mysqlRouteGroup.GET("/delete", handlers.MysqlDelete)


		mysqlRouteGroup.GET("/orm-one", handlers.MysqlSelectOneOrm)
	}


	mysqlOrmRouteGroup := router.Group("/rnd/mysql-orm")
	{
		mysqlOrmRouteGroup.GET("/one", handlers.MysqlSelectOneOrm)

	}
}