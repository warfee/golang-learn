package handlers

import (
	"fmt"
  	"os"
  	"net/http"
  	"math/rand"

  	"time"
  	"context"
  	"github.com/gin-gonic/gin"
    "go.mongodb.org/mongo-driver/v2/mongo"
    "go.mongodb.org/mongo-driver/v2/bson"
    "go.mongodb.org/mongo-driver/v2/mongo/options"
)

const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

func ConnectMongo() (*mongo.Client, error) {

	mongoURL := os.Getenv("MONGO_URL")
	if mongoURL == "" {
		return nil, fmt.Errorf("MONGO_URL not set")
	}

	opts := options.Client().SetTimeout(5 * time.Second)
	client, err := mongo.Connect(opts.ApplyURI(mongoURL))

	if err != nil {
		return nil, err
	}

	return client,nil
}

func MongoSelectOne(c *gin.Context) {

	client, err := ConnectMongo()

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error" : true,
			"message": "MongoDB connection failed", 
			"details": err.Error(),
		})
		return
	}

	collection := client.Database("warfee").Collection("warfee-testing")

	var result bson.M
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	err = collection.FindOne(ctx, bson.D{}).Decode(&result)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			c.JSON(http.StatusNotFound, gin.H{
				"error" : true,
				"message": "No document found",
			})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{
			"error" : true,
			"message": "Find failed", 
			"details": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status": "success",
		"data":   result,
	})
}

func MongoInsert(c *gin.Context) {

	client, err := ConnectMongo()

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error" : true,
			"message": "MongoDB connection failed", 
			"details": err.Error(),
		})
		return
	}

	collection := client.Database("warfee").Collection("warfee-testing")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	doc := bson.M{
		"name":       "Testing",
		"email":      "test@test.com",
		"membership": "gold",
		"points":     1500,
		"created_at": time.Now(),
	}

	result, err := collection.InsertOne(ctx, doc)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error" : true,
			"message": "Insertion failed", 
			"details": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"error": false,
		"response_result": result,
	})

}

func MongoUpdate(c *gin.Context) {

	client, err := ConnectMongo()

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error" : true,
			"message": "MongoDB connection failed", 
			"details": err.Error(),
		})
		return
	}

	collection := client.Database("warfee").Collection("warfee-testing")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()


	findOpts := options.FindOne().SetSort(bson.D{{"created_at", -1}})

	var latestDoc bson.M

	err = collection.FindOne(ctx, bson.D{}, findOpts).Decode(&latestDoc)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error" : true,
			"message": "No latest record found", 
			"details": err.Error(),
		})
		return
	}

	randomString := RandomString(10)

    filter := bson.M{"_id": latestDoc["_id"]} // filter
	update := bson.M{"$set": bson.M{"name": randomString}}

	res, err := collection.UpdateOne(ctx, filter,update)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error" : true,
			"message": "Deletion failed", 
			"details": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"error": false,
		"response_result": res,
		"updated_name" : randomString,
	})

}

func MongoDelete(c *gin.Context) {

	client, err := ConnectMongo()

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error" : true,
			"message": "MongoDB connection failed", 
			"details": err.Error(),
		})
		return
	}

	collection := client.Database("warfee").Collection("warfee-testing")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()


	findOpts := options.FindOne().SetSort(bson.D{{"created_at", -1}})

	var latestDoc bson.M

	err = collection.FindOne(ctx, bson.D{}, findOpts).Decode(&latestDoc)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error" : true,
			"message": "No latest record found", 
			"details": err.Error(),
		})
		return
	}

    filter := bson.M{"_id": latestDoc["_id"]}
	res, err := collection.DeleteOne(ctx, filter)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error" : true,
			"message": "Deletion failed", 
			"details": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"error": false,
		"response_result": res,
	})

}

func RandomString(length int) string {
	rand.Seed(time.Now().UnixNano()) // seed only once
	b := make([]byte, length)
	for i := range b {
		b[i] = charset[rand.Intn(len(charset))]
	}
	return string(b)
}





