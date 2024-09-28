package models

import (
	"imta.icu/rest/db"
	"log"
	"time"
)

type Event struct {
	ID          int64
	Name        string `binding:"required"`
	Description string `binding:"required"`
	Location    string `binding:"required"`
	Date        time.Time
	UserId      int
}

var events []Event

func (event Event) Save() error {
	// TODO: Add to data base

	query := `
			INSERT INTO events(name, description, location, dateTime, user_id) 
			VALUES (?, ?, ?, ?, ?)`
	stmt, err := db.DB.Prepare(query)
	if err != nil {
		log.Fatal(err)
		return err
	}

	defer stmt.Close()

	result, err := stmt.Exec(event.Name, event.Description, event.Location, event.Date, event.UserId)
	if err != nil {
		log.Fatal(err)
		return err
	}

	id, err := result.LastInsertId()
	if err != nil {
		log.Fatal(err)
		return err
	}

	event.ID = id

	return err

}

func GetAllEvents() ([]Event, error) {

	query := "SELECT * FROM events"
	rows, err := db.DB.Query(query)
	if err != nil {
		log.Fatal(err)
		return nil, err
	}

	var events []Event
	for rows.Next() {
		var event Event
		err := rows.Scan(&event.ID, &event.Name, &event.Description, &event.Location, &event.Date, &event.UserId)
		if err != nil {
			log.Fatal(err)
			return nil, err
		}

		events = append(events, event)
	}

	return events, nil
}
