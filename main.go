package main

import (
  "log"
  "os"

  "github.com/gin-gonic/gin"
  "github.com/joho/godotenv"
  "github.com/hibiken/asynq"
  "github.com/hibiken/asynqmon"

  "proj1/routes"
  "proj1/tasks"
)

const redisAddr = "165.232.172.114:6379"

func main() {
  // Load env
  err := godotenv.Load()
  if err != nil {
    log.Println("No .env file found, using default port")
  }

  // Start Asynq worker
  srv := asynq.NewServer(
    asynq.RedisClientOpt{Addr: redisAddr},
    asynq.Config{Concurrency: 10},
  )

  mux := asynq.NewServeMux()
  mux.HandleFunc(tasks.TypeSlackSendMessage, tasks.HandleSlackSendMessageTask)

  go func() {
    if err := srv.Run(mux); err != nil {
      log.Fatalf("could not run Asynq server: %v", err)
    }
  }()

  // Setup monitoring
  h := asynqmon.New(asynqmon.Options{
    RootPath:     "/monitoring",
    RedisConnOpt: asynq.RedisClientOpt{Addr: redisAddr},
  })

  // Start Gin
  port := os.Getenv("PORT")
  if port == "" {
    port = "8080"
  }

  router := gin.Default()
  routes.SetupRoutes(router)
  router.Any("/monitoring/*path", gin.WrapH(h))

  err = router.Run(":" + port)
  if err != nil {
    log.Fatalf("could not run Gin server: %v", err)
  }
}