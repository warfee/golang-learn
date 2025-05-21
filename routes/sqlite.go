package routes

import (
	"github.com/gin-gonic/gin"
	"proj1/handlers"
)

func SqliteRoutes(router *gin.Engine) {
	sqliteRouteGroup := router.Group("/sqlite")
	{
		sqliteRouteGroup.GET("/check", handlers.SQLiteCheck)
		sqliteRouteGroup.GET("/migrate", handlers.SQLiteMigration)
		sqliteRouteGroup.GET("/insert", handlers.SQLiteInsert)
		sqliteRouteGroup.GET("/all", handlers.SQLiteAll)


		sqliteRouteGroup.GET("/all-monitor", handlers.SQLiteAllMonitor)
	}
}