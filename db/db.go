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
    user_id INTEGER NOT NULL
)`

	_, err := DB.Exec(createEventsTable)

	if err != nil {
		fmt.Println("ERROR ", err)
		panic("failed to create events table")
	}

	_, errr := DB.Exec(createActionsTable)
	if errr != nil {
		fmt.Println("ERROR ", errr)
		panic("failed to create events table")
	}

}
