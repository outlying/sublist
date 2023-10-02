package main

import (
	"database/sql"
	"log"

	"github.com/go-jet/jet/v2/sqlite"
	_ "github.com/mattn/go-sqlite3"
)

type User struct {
	Email    string `jet:"type:TEXT; PRIMARY KEY"`
	Name     string `jet:"type:TEXT"`
	Password string `jet:"type:TEXT"`
}

type ToDoList struct {
	ID        int64  `jet:"type:INTEGER; PRIMARY KEY AUTOINCREMENT"`
	Name      string `jet:"type:TEXT"`
	UserEmail string `jet:"type:TEXT"`
}

type Item struct {
	ID        int64  `jet:"type:INTEGER; PRIMARY KEY AUTOINCREMENT"`
	Name      string `jet:"type:TEXT"`
	ListID    int64  `jet:"type:INTEGER"`
	SubListID int64  `jet:"type:INTEGER"`
}

func setupDatabase() *sql.DB {
	// Open SQLite3 database
	db, err := sql.Open("sqlite3", "./test.db")
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}

	// Create User table
	err = sqlite.CREATE_TABLE(User{}).Exec(db)
	if err != nil {
		log.Fatal("Failed to create User table:", err)
	}

	// Create ToDoList table
	err = sqlite.CREATE_TABLE(ToDoList{}).Exec(db)
	if err != nil {
		log.Fatal("Failed to create ToDoList table:", err)
	}

	// Create Item table
	err = sqlite.CREATE_TABLE(Item{}).Exec(db)
	if err != nil {
		log.Fatal("Failed to create Item table:", err)
	}

	return db
}
