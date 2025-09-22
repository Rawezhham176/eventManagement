package model_test

import (
	"database/sql"
	"eventManagement/db"
	"eventManagement/model"
	"testing"
	"time"
)

func setupTestEventDB(t *testing.T) *sql.DB {
	t.Helper()

	testDB, err := sql.Open("sqlite3", ":memory:")
	if err != nil {
		t.Fatalf("open in-memory db: %v", err)
	}

	schema := `
	CREATE TABLE events (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		name TEXT NOT NULL,
		description TEXT NOT NULL,
		location TEXT NOT NULL,
		dateTime DATETIME NOT NULL,
		user_id INTEGER
	);`
	if _, err := testDB.Exec(schema); err != nil {
		t.Fatalf("create schema: %v", err)
	}

	now := time.Now().UTC().Format(time.RFC3339)
	insert := `INSERT INTO events (name, description, location, dateTime, user_id) VALUES (?, ?, ?, ?, ?);`
	if _, err := testDB.Exec(insert, "Meetup", "Go User Group", "Somewhere", now, 11); err != nil {
		t.Fatalf("insert row 1: %v", err)
	}
	if _, err := testDB.Exec(insert, "Conference", "Dev Summit", "Berlin", now, 22); err != nil {
		t.Fatalf("insert row 2: %v", err)
	}

	return testDB
}

func TestGetAllEvents_ReturnsRows(t *testing.T) {
	testDB := setupTestEventDB(t)
	defer testDB.Close()

	db.DB = testDB

	events, err := model.GetAllEvents()
	if err != nil {
		t.Fatalf("GetAllEvents error: %v", err)
	}
	if len(events) != 2 {
		t.Fatalf("expected 2 event, got %d", len(events))
	}

	if events[0].Name == "" || events[0].Location == "" {
		t.Errorf("unexpected empty fields: %+v", events[0])
	}
	if events[1].UserID == 0 {
		t.Errorf("expected non-zero UserID, got %d", events[1].UserID)
	}
}
