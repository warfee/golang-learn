package handlers

import (
  "log"
  "database/sql"
    "github.com/gin-gonic/gin"
  _ "github.com/ncruces/go-sqlite3/driver"
  _ "github.com/ncruces/go-sqlite3/embed"
)

func SQLiteCheck(c *gin.Context) {

  var version string
  db, _ := sql.Open("sqlite3", "database.sqlite")
  db.QueryRow(`SELECT sqlite_version()`).Scan(&version)

	c.JSON(200, gin.H{
      "version": version,
    })
}

func SQLiteMigration(c *gin.Context) {
  db, err := sql.Open("sqlite3", "database.sqlite")
  if err != nil {
    log.Printf("Error opening sqlite file: %v", err)
    c.JSON(500, gin.H{"error": "failed to open database"})
    return
  }
  defer db.Close()

  _, err = db.Exec(`CREATE TABLE IF NOT EXISTS users (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    name TEXT
    category_id TEXT
  )`)


  _, err = db.Exec(`CREATE TABLE IF NOT EXISTS monitor (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    cpu float64,
    memory float64,
    storage float64
  )`)

  if err != nil {
    log.Printf("Error migrating database: %v", err)
    c.JSON(500, gin.H{"error": "failed to run migration"})
    return
  }

  c.JSON(200, gin.H{
    "done": true,
  })
}

func SQLiteInsert(c *gin.Context){

  db, err := sql.Open("sqlite3", "database.sqlite")

  if err != nil {

    log.Fatal(err)
  }
  defer db.Close()

  res, err := db.Exec(`INSERT INTO users (name) VALUES (?)`, "John Doe")
  if err != nil {
    log.Fatal(err)
  }

  // Optional: Get last inserted ID
  id, _ := res.LastInsertId()
  log.Printf("Inserted user with ID: %d\n", id)

}


func SQLiteAll(c *gin.Context) {
  db, err := sql.Open("sqlite3", "database.sqlite")
  if err != nil {
    log.Fatal(err)
  }
  defer db.Close()

  rows, err := db.Query(`SELECT id, name FROM users`)
  if err != nil {
    log.Fatal(err)
  }
  defer rows.Close()

  var users []map[string]interface{}

  for rows.Next() {
    var id int
    var name string

    err := rows.Scan(&id, &name)
    if err != nil {
      log.Fatal(err)
    }

    users = append(users, gin.H{
      "id":          id,
      "name":        name,
    })
  }

  c.JSON(200, users)
}

func SQLiteAllMonitor(c *gin.Context) {
  db, err := sql.Open("sqlite3", "database.sqlite")
  if err != nil {
    log.Fatal(err)
  }
  defer db.Close()

  rows, err := db.Query(`SELECT id, cpu, memory, storage FROM monitor`)
  if err != nil {
    log.Fatal(err)
  }
  defer rows.Close()

  var users []map[string]interface{}

  for rows.Next() {
    var id int
    var cpu string
    var memory string
    var storage string

    err := rows.Scan(&id, &cpu,&memory,&storage)
    if err != nil {
      log.Fatal(err)
    }

    users = append(users, gin.H{
      "id":          id,
      "cpu":        cpu,
      "memory":        memory,
      "storage":        storage,
    })
  }

  c.JSON(200, users)
}






