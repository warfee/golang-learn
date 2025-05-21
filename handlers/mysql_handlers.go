package handlers

import (
	"database/sql"
	"fmt"
	"net/http"
	"os"
	"time"
	"math/rand"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
)


const mysql_charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

func MysqlRandomString(length int) string {
	rand.Seed(time.Now().UnixNano()) // seed only once
	b := make([]byte, length)
	for i := range b {
		b[i] = charset[rand.Intn(len(mysql_charset))]
	}
	return string(b)
}


func MysqlConnect() (*sql.DB, error) {

	host := os.Getenv("MYSQL_DB_HOST")
	port := os.Getenv("MYSQL_DB_PORT")
	name := os.Getenv("MYSQL_DB_NAME")
	username := os.Getenv("MYSQL_DB_USERNAME")
	password := os.Getenv("MYSQL_DB_PASSWORD")

	if host == "" || port == "" || name == "" || username == "" || password == "" {

		return nil, fmt.Errorf("database connection info is missing")
	}

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", username, password, host, port, name)

	db, err := sql.Open("mysql", dsn)

	if err != nil {
		return nil, fmt.Errorf("failed to open connection: %v", err)
	}

	if err = db.Ping(); err != nil {
		return nil, fmt.Errorf("ping failed: %v", err)
	}

	return db, nil
}


func MysqlSelectOne(c *gin.Context) {

	db, err := MysqlConnect()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error":   true,
			"message": "Mysql connection failed",
			"details": err.Error(),
		})
		return
	}
	defer db.Close()

	var id int
	var name string
	var category int
	var created_at string

	if err := db.QueryRow("SELECT id, name, category_id ,created_at FROM `users` LIMIT 1").Scan(&id, & name,&category, &created_at); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error":   true,
			"message": "Failed to query",
			"details": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"id": id,
		"name": name,
		"category": category,
		"created_at": created_at,
	})
}


func MysqlInsert(c *gin.Context) {

	db, err := MysqlConnect()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error":   true,
			"message": "Mysql connection failed",
			"details": err.Error(),
		})
		return
	}
	defer db.Close()

	randomString := MysqlRandomString(10)

	name := "User " + randomString
	category_id := 1
	created_at := time.Now()

	query := `INSERT INTO users (name, category_id, created_at) VALUES (?, ?, ?)`

	result, err := db.Exec(query, name, category_id, created_at)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error":   true,
			"message": "Insert failed",
			"details": err.Error(),
		})
		return
	}

	insertedID, _ := result.LastInsertId()

	c.JSON(http.StatusOK, gin.H{
		"message":      "Insertion successful",
		"inserted_id":  insertedID,
	})
}

func MysqlUpdate(c *gin.Context) {

	db, err := MysqlConnect()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error":   true,
			"message": "Mysql connection failed",
			"details": err.Error(),
		})
		return
	}
	defer db.Close()

	var id int

	if err := db.QueryRow("SELECT id FROM `users` LIMIT 1").Scan(&id); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error":   true,
			"message": "Failed to query",
			"details": err.Error(),
		})
		return
	}

	randomString := MysqlRandomString(10)
	name := "User " + randomString 

	query := `UPDATE users SET name = ? WHERE id = ?`

	result, err := db.Exec(query, name, id)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error":   true,
			"message": "Update record failed",
			"details": err.Error(),
		})
		return
	}

	rowsAffected, _ := result.RowsAffected()

	c.JSON(http.StatusOK, gin.H{
		"id":  id,
		"row_affected" : rowsAffected,
		"name" : name,
	})
}

func MysqlDelete(c *gin.Context) {

	db, err := MysqlConnect()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error":   true,
			"message": "Mysql connection failed",
			"details": err.Error(),
		})
		return
	}
	defer db.Close()

	var id int

	if err := db.QueryRow("SELECT id FROM `users` LIMIT 1").Scan(&id); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error":   true,
			"message": "Failed to query",
			"details": err.Error(),
		})
		return
	}

	query := `DELETE FROM users WHERE id = ?`

	result, err := db.Exec(query, id)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error":   true,
			"message": "Update record failed",
			"details": err.Error(),
		})
		return
	}

	rowsAffected, _ := result.RowsAffected()

	c.JSON(http.StatusOK, gin.H{
		"id":  id,
		"row_affected" : rowsAffected,
	})
}










