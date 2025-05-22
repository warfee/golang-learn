package handlers

import (
	"os"
	"github.com/gin-gonic/gin"
    "log"
    "net/http"
    "github.com/hibiken/asynq"
    "proj1/tasks"
)

const redisAddr = "165.232.172.114:6379"

func QueueSlack(c *gin.Context) {
	client := asynq.NewClient(asynq.RedisClientOpt{Addr: redisAddr})
	defer client.Close()


	SlackWebHook := os.Getenv("SLACK_WEBHOOK")

	task, err := tasks.NewSlackSendMessageTask(SlackWebHook, "warfee test message slack")
	if err != nil {
		log.Printf("could not create task: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create task"})
		return
	}

	info, err := client.Enqueue(task)
	if err != nil {
		log.Printf("could not enqueue task: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to enqueue task"})
		return
	}

	log.Printf("enqueued task: id=%s queue=%s", info.ID, info.Queue)
	c.JSON(http.StatusOK, gin.H{
		"message": "Task enqueued",
		"task_id": info.ID,
		"queue":   info.Queue,
	})
}