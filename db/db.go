package db

import (
	"database/sql"
)

var DB *sql.DB

func InitDB() {

	DB, err := sql.Open("sqllite3", "api.db")

	if err != nil {
		panic("Something went wromg with db**")
	}

}
