package db

import (
	"database/sql"

	_ "github.com/mattn/go-sqlite3"
)

var DB *sql.DB

func InitDatabase() {

	var err error
	DB, err = sql.Open("sqlite3", "api.db")

	if err != nil {
		panic("Could not create connection to database!")
	}

	DB.SetMaxOpenConns(10)
	DB.SetMaxIdleConns(5)

	createTables()
}

func createTables() {
	createNotesTable :=
		`CREATE TABLE IF NOT EXISTS notes (
			id INTEGER PRIMARY KEY AUTOINCREMENT, 
			title TEXT,
			description TEXT,
			creationDate DATETIME NOT NULL,
			userId INTEGER,
			FOREIGN KEY (userId) REFERENCES users(id)
			)`

	_, err := DB.Exec(createNotesTable)

	if err != nil {
		panic("Creating notes table was unsuccesfull!")
	}

	createUsersTable :=
		`CREATE TABLE IF NOT EXISTS users (
			id INTEGER PRIMARY KEY AUTOINCREMENT, 
			email TEXT NOT NULL UNIQUE,
			password TEXT NOT NULL
			)`

	_, err = DB.Exec(createUsersTable)

	if err != nil {
		panic("Creating users table was unsuccesfull!")
	}
}
