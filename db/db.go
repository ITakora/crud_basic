package db

import (
	"database/sql"

	_ "github.com/mattn/go-sqlite3"
)

var DB *sql.DB

func InitDB() {
	var err error
	DB, err = sql.Open("sqlite3", "user.db")

	if err != nil {
		panic("Something went wrong with db connection")
	}

	DB.SetMaxIdleConns(5)
	DB.SetMaxOpenConns(10)
	createTable()
}

func createTable() {
	createUsersTable := `CREATE TABLE IF NOT EXISTS users (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		name TEXT NOT NULL,
		age INTEGER NOT NULL,
		deskription TEXT NOT NULL
	)`

	_, err := DB.Exec(createUsersTable)

	if err != nil {
		panic("Could not create users table")
	}
}
