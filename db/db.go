package db

import (
	"database/sql"
	"fmt"

	_ "github.com/mattn/go-sqlite3"
)

var DB *sql.DB

func InitDB() {
	var err error
	DB, err = sql.Open("sqlite3", "api.db")
	if err != nil {
		panic(fmt.Sprintf("open db error: %s", err))
	}

	DB.SetMaxOpenConns(10)
	DB.SetMaxIdleConns(5)
	createTables()
}

func createTables() {
	createUserTable := `
       CREATE TABLE IF NOT EXISTS users (
           id INTEGER PRIMARY KEY NOT NULL AUTOINCREMENT,
           email TEXT NOT NULL,
           password TEXT NOT NULL,
       )`

	_, err := DB.Exec(createUserTable)
	if err != nil {
		panic(fmt.Sprintf("create events table error: %s", err))
	}

	createEventsTable := `
       CREATE TABLE IF NOT EXISTS events (
           id INTEGER PRIMARY KEY AUTOINCREMENT,
           name TEXT NOT NULL,
           description TEXT NOT NULL,
           location TEXT NOT NULL,
           dateTime DATETIME NOT NULL,
           user_id INTEGER,
           FOREIGN KEY(user_id) REFERENCES users(id)
       )`

	_, err = DB.Exec(createEventsTable)
	if err != nil {
		panic(fmt.Sprintf("create events table error: %s", err))
	}

	createRegistrationTable := `
       CREATE TABLE IF NOT EXISTS registration (
           id INTEGER PRIMARY KEY AUTOINCREMENT,
           user_id INTEGER,
           event_id INTEGER,
           FOREIGN KEY(event_id) REFERENCES events(id),
           FOREIGN KEY(user_id) REFERENCES users(id),
       )`

	_, err = DB.Exec(createRegistrationTable)
	if err != nil {
		panic(fmt.Sprintf("create registration table error: %s", err))
	}
}
