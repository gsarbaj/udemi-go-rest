package db

import (
	"database/sql"
	"fmt"
	_ "github.com/mattn/go-sqlite3"
)

var DB *sql.DB

func InitDb() {
	var err error
	DB, err = sql.Open("sqlite3", "api.db")
	if err != nil {
		fmt.Println("ERROR ", err)
		panic("failed to connect database")
	}
	DB.SetMaxOpenConns(10)
	DB.SetMaxIdleConns(5)

	createTables()
}

func createTables() {

	createUserTable := `CREATE TABLE IF NOT EXISTS users (
    	id INTEGER PRIMARY KEY AUTOINCREMENT,
    	email VARCHAR(255) UNIQUE NOT NULL,
    	password VARCHAR(255) NOT NULL
)`

	createEventsTable := `CREATE TABLE IF NOT EXISTS events (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    name TEXT NOT NULL,
    description TEXT NOT NULL,
    location TEXT NOT NULL,
    dateTime DateTime NOT NULL,
    user_id INTEGER NOT NULL
)`

	createActionsTable := `CREATE TABLE IF NOT EXISTS actions (
     id INTEGER PRIMARY KEY AUTOINCREMENT,
    name TEXT NOT NULL,
    description TEXT NOT NULL,
    location TEXT NOT NULL,
    dateTime DateTime NOT NULL,
    user_id INTEGER NOT NULL,
	FOREIGN KEY(user_id) REFERENCES users(id)
)`

	_, err := DB.Exec(createEventsTable)

	if err != nil {
		fmt.Println("ERROR ", err)
		panic("failed to create events table")
	}

	_, err = DB.Exec(createActionsTable)
	if err != nil {
		fmt.Println("ERROR ", err)
		panic("failed to create action table")
	}

	_, err = DB.Exec(createUserTable)
	if err != nil {
		fmt.Println("ERROR ", err)
		panic("failed to create users table")
	}

}
