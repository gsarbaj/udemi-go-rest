package models

import (
	"database/sql"
	"imta.icu/rest/db"
	"time"
)

type Action struct {
	ID          int64
	Name        string    `binding:"required"`
	Description string    `binding:"required"`
	Location    string    `binding:"required"`
	DateTime    time.Time `binding:"required"`
	UserID      int
}

var actions []Action = []Action{}

func (action Action) Save() error {
	// TODO: add action to a database

	query := `INSERT INTO actions(name, description, location, dateTime, user_id)
			  VALUES (?, ?, ?, ?, ?)`

	stmt, err := db.DB.Prepare(query)
	if err != nil {
		return err
	}
	defer stmt.Close()

	result, err := stmt.Exec(action.Name, action.Description, action.Location, action.DateTime, action.UserID)
	if err != nil {
		return err
	}

	id, err := result.LastInsertId()
	action.ID = id

	return err

}

func GetAllActions() ([]Action, error) {
	query := `SELECT * FROM actions`
	rows, err := db.DB.Query(query)
	if err != nil {
		return nil, err
	}
	defer func(rows *sql.Rows) {
		err := rows.Close()
		if err != nil {
		}
	}(rows)

	var actions []Action
	for rows.Next() {
		var action Action
		err := rows.Scan(&action.ID, &action.Name, &action.Description, &action.Location, &action.DateTime, &action.UserID)
		if err != nil {
			return nil, err
		}
		actions = append(actions, action)
	}

	return actions, nil
}

func GetActionByID(id int64) (*Action, error) {
	query := `SELECT * FROM actions WHERE id = ?`
	row := db.DB.QueryRow(query, id)

	var action Action
	err := row.Scan(&action.ID, &action.Name, &action.Description, &action.Location, &action.DateTime, &action.UserID)
	if err != nil {
		return nil, err
	}
	return &action, nil
}

func (action Action) UpdateAction() error {
	query := `
UPDATE actions
SET name = ?, description = ?, location = ?, dateTime = ?
WHERE id = ?
`
	stmt, err := db.DB.Prepare(query)
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(action.Name, action.Description, action.Location, action.DateTime, action.ID)
	return err
}

func (action Action) DeleteAction() error {
	query := `DELETE FROM actions WHERE id = ?`
	stmt, err := db.DB.Prepare(query)
	if err != nil {
		return err
	}
	defer stmt.Close()
	_, err = stmt.Exec(action.ID)
	return err
}
