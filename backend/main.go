package main

import (
	"database/sql"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	_ "github.com/mattn/go-sqlite3"
)

var db *sql.DB

func setupDatabase() {
	var err error
	db, err = sql.Open("sqlite3", "./todo.db")
	if err != nil {
		log.Fatal(err)
	}

	statements := []string{
		`CREATE TABLE IF NOT EXISTS users (id INTEGER PRIMARY KEY AUTOINCREMENT, username TEXT, password TEXT);`,
		`CREATE TABLE IF NOT EXISTS lists (id INTEGER PRIMARY KEY AUTOINCREMENT, title TEXT, user_id INTEGER, FOREIGN KEY(user_id) REFERENCES users(id));`,
	}

	for _, stmt := range statements {
		_, err = db.Exec(stmt)
		if err != nil {
			log.Fatalf("Failed to execute statement: %v", err)
		}
	}
}

func main() {
	setupDatabase()
	r := gin.Default()

	// Register
	r.POST("/register", func(c *gin.Context) {
		username := c.PostForm("username")
		password := c.PostForm("password")

		_, err := db.Exec("INSERT INTO users (username, password) VALUES (?, ?)", username, password)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not register user"})
			return
		}
		c.JSON(http.StatusOK, gin.H{"message": "User registered"})
	})

	// Create list (Assume user_id is 1)
	r.POST("/lists", func(c *gin.Context) {
		title := c.PostForm("title")

		_, err := db.Exec("INSERT INTO lists (title, user_id) VALUES (?, ?)", title, 1)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not create list"})
			return
		}
		c.JSON(http.StatusOK, gin.H{"message": "List created"})
	})

	// Get lists (Assume user_id is 1)
	r.GET("/lists", func(c *gin.Context) {
		rows, err := db.Query("SELECT id, title FROM lists WHERE user_id = ?", 1)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not fetch lists"})
			return
		}
		defer rows.Close()

		lists := []gin.H{}
		for rows.Next() {
			var id int
			var title string
			rows.Scan(&id, &title)
			lists = append(lists, gin.H{"id": id, "title": title})
		}
		c.JSON(http.StatusOK, gin.H{"lists": lists})
	})

	r.Run(":8080")
}
