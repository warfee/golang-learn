package main


import (
  "log"
  "os"

  "github.com/gin-gonic/gin"
  "github.com/joho/godotenv"
  "proj1/routes"
  "proj1/handlers"
)


func main() {

  err := godotenv.Load()
  if err != nil {
    log.Println("No .env file found, using default port")
  }

  port := os.Getenv("PORT")
  if port == "" {
    port = "8080"
  }

  handlers.StartMetricsTicker()

  router := gin.Default()

  routes.SetupRoutes(router)

  err = router.Run(":" + port)  
}


