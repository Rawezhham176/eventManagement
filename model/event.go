package model

import (
	"database/sql"
	"eventManagement/db"
	"fmt"
	"time"
)

type Event struct {
	ID          int64     `json:"id"`
	Name        string    `json:"name" binding:"required"`
	Description string    `json:"description" binding:"required"`
	Location    string    `json:"location" binding:"required"`
	Category    string    `json:"category" binding:"required"`
	DateTime    time.Time `json:"dateTime" binding:"required"`
	UserID      int64     `json:"user_id"`
}

func GetAllEvents() ([]Event, error) {
	query := `SELECT * FROM events`
	rows, err := db.DB.Query(query)
	if err != nil {
		return nil, err
	}
	defer func(rows *sql.Rows) {
		err := rows.Close()
		if err != nil {
			panic(fmt.Sprintf("close events table error: %s", err))
		}
	}(rows)

	var events []Event

	for rows.Next() {
		var event Event
		err := rows.Scan(&event.ID, &event.Name, &event.Description, &event.Location, &event.DateTime, &event.UserID)
		if err != nil {
			return nil, err
		}
		events = append(events, event)
	}

	return events, nil
}

func GetEventById(id int64) (*Event, error) {
	query := `SELECT * FROM events WHERE id = ?`
	row := db.DB.QueryRow(query, id)
	var event Event
	err := row.Scan(&event.ID, &event.Name, &event.Description, &event.Location, &event.DateTime, &event.UserID)
	if err != nil {
		return nil, err
	}
	return &event, nil
}

func GetEventByNamOrLocation(name, location string) (*Event, error) {
	query := `SELECT * FROM events WHERE name = ? OR location = ?`
	row := db.DB.QueryRow(query, name, location)
	var event Event
	err := row.Scan(&event.ID, &event.Name, &event.Description, &event.Location, &event.DateTime, &event.UserID)
	if err != nil {
		return nil, err
	}
	return &event, nil
}

func GetEventsByCategory(category string) (*Event, error) {
	query := `SELECT * FROM events WHERE category = ?`
	row := db.DB.QueryRow(query, category)
	var event Event
	err := row.Scan(&event.ID, &event.Name, &event.Description, &event.Location, &event.Category, &event.DateTime, &event.UserID)
	if err != nil {
		return nil, err
	}
	return &event, nil
}

func GetEventsByUserId(userId int64) (*[]Event, error) {
	query := `SELECT * FROM events WHERE user_id = ?`
	rows, err := db.DB.Query(query, userId)
	if err != nil {
		return nil, err
	}
	var events []Event
	for rows.Next() {
		var event Event
		err := rows.Scan(&event.ID, &event.Name, &event.Description, &event.Location, &event.DateTime, &event.UserID)
		if err != nil {
			return nil, err
		}
		events = append(events, event)
	}
	return &events, nil
}

func GetUpcomingEvents(startDate, endDate string) ([]Event, error) {
	query := `SELECT * FROM events WHERE dateTime BETWEEN ? AND ?`
	rows, err := db.DB.Query(query, startDate, endDate)

	defer func(rows *sql.Rows) {
		err := rows.Close()
		if err != nil {
			return
		}
	}(rows)

	if err != nil {
		return nil, err
	}

	var events []Event

	for rows.Next() {
		var event Event
		err := rows.Scan(&event.ID, &event.Name, &event.Description, &event.Location, &event.Category, &event.DateTime, &event.UserID)
		if err != nil {
			return nil, err
		}
		events = append(events, event)
	}
	return events, nil
}

func (e *Event) Save() error {
	query := `
		INSERT INTO events(name, description, location,dateTime,user_id)
		VALUES(?,?,?,?,?)`
	prepare, err := db.DB.Prepare(query)
	if err != nil {
		panic(fmt.Sprintf("create events table error: %s", err))
	}
	result, err := prepare.Exec(e.Name, e.Description, e.Location, e.DateTime, e.UserID)
	if err != nil {
		panic(fmt.Sprintf("create events table error: %s", err))
	}
	defer func(prepare *sql.Stmt) {
		err := prepare.Close()
		if err != nil {
			panic(fmt.Sprintf("close events table error: %s", err))
		}
	}(prepare)

	id, err := result.LastInsertId()
	if err != nil {
		panic(fmt.Sprintf("create events table error: %s", err))
	}
	e.ID = id
	return err
}

func (e Event) UpdateEvent() error {
	query := `
	UPDATE events
    SET name = ?, description = ?, location = ?, dateTime = ?, user_is = ?
	WHERE  id = ?
`
	prepare, err := db.DB.Prepare(query)

	if err != nil {
		panic(fmt.Sprintf("prepare update events table error: %s", err))
	}

	defer func(prepare *sql.Stmt) {
		err := prepare.Close()
		if err != nil {
			panic(fmt.Sprintf("close events table error: %s", err))
		}
		return
	}(prepare)

	_, err = prepare.Exec(prepare)
	return err
}

func (e Event) DeleteEvent() error {
	query := `DELETE FROM events WHERE  id = ?`
	prepare, err := db.DB.Prepare(query)

	if err != nil {
		panic(fmt.Sprintf("prepare update events table error: %s", err))
	}

	defer func(prepare *sql.Stmt) {
		err := prepare.Close()
		if err != nil {
			panic(fmt.Sprintf("close events table error: %s", err))
		}
		return
	}(prepare)

	_, err = prepare.Exec(e.ID)
	return err
}

func (e Event) RegisterEvent(userID int64) error {
	query := `INSERT INTO registrations(event_id, user_id) VALUES(?,?)`
	stmt, err := db.DB.Prepare(query)
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(e.ID, userID)
	return err
}

func (e Event) CancelEvent(userID int64) error {
	query := `DELETE FROM registrations WHERE event_id = ? AND user_id = ?`
	stmt, err := db.DB.Prepare(query)
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(e.ID, userID)
	return err
}

func (e Event) GetRegistrationList(eventID int64) ([]Registration, error) {
	query := `SELECT * FROM registrations WHERE event_id = ?`
	stmt, err := db.DB.Prepare(query)
	if err != nil {
		return nil, err
	}
	defer func(stmt *sql.Stmt) {
		err := stmt.Close()
		if err != nil {
			return
		}
	}(stmt)

	row, err := stmt.Query(eventID)
	if err != nil {
		return nil, err
	}

	var registrationList []Registration

	for row.Next() {
		var registration Registration
		err := row.Scan(&registration.ID, &registration.UserID, &registration.EventID)
		if err != nil {
			return nil, err
		}
		registrationList = append(registrationList, registration)
	}

	return registrationList, nil
}

func GetRegistrationsByUserId(userID int64) ([]Registration, error) {
	query := `SELECT FROM registrations WHERE user_id = ?`
	rows, err := db.DB.Query(query, userID)
	if err != nil {
		return nil, err
	}
	var registrationList []Registration
	for rows.Next() {
		var registration Registration
		err := rows.Scan(&registration.ID, &registration.UserID, &registration.EventID)
		if err != nil {
			return nil, err
		}
		registrationList = append(registrationList, registration)
	}
	return registrationList, nil
}
