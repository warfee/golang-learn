package routes

import (
	"github.com/gin-gonic/gin"
)

func SetupRoutes(router *gin.Engine) {

	RndRoutes(router)
	MongoRoutes(router)
	MysqlRoutes(router)
}