package sqlserver

import (
	"context"
	"database/sql"
	"fmt"
	"github.com/vikashvverma/eventers/pkg/utl/model"
)

// NewEvent returns a new event database instance
func NewEvent() *Event {
	return &Event{}
}

// Event represents the client for user table
type Event struct{}

// Create creates a new event on database
func (u *Event) Create(db *sql.DB, e eventers.Event) (int64, error) {
	ctx := context.Background()
	var err error

	// Check if database is alive.
	err = db.PingContext(ctx)
	if err != nil {
		return -1, err
	}

	tsql := `INSERT INTO Event (Name, Date, Location) VALUES (@Name, @Date, @Location); SELECT convert(bigint, SCOPE_IDENTITY());`

	// Execute query
	stmt, err := db.Prepare(tsql)
	if err != nil {
		return -1, err
	}
	defer stmt.Close()

	row := stmt.QueryRowContext(
		ctx,
		sql.Named("Name", e.Name),
		sql.Named("Date", e.Date),
		sql.Named("Location", e.Location),
	)

	if row == nil {
		return -1, fmt.Errorf("invalid data")
	}
	var newID int64
	err = row.Scan(&newID)
	if err != nil {
		return -1, fmt.Errorf("unable to create event, check data")
	}

	return newID, nil
}

// View returns single event by ID
func (u *Event) View(db *sql.DB, id int) (*eventers.Event, error) {
	ctx := context.Background()

	// Check if database is alive.
	err := db.PingContext(ctx)
	if err != nil {
		return nil, err
	}

	tsql := fmt.Sprintf("SELECT Name, Date, Location FROM Event WHERE ID=@ID;")

	// Execute query
	stmt, err := db.Prepare(tsql)
	if err != nil {
		return nil, err
	}
	defer stmt.Close()

	row := stmt.QueryRow(
		ctx,
		sql.Named("ID", id),
	)
	if err != nil {
		return nil, err
	}

	e := eventers.Event{}

	// if row exists.
	err = row.Scan(&e.Name, &e.Date, &e.Location)
	if err != nil {
		return nil, err
	}

	return &e, nil
}

// Update updates event's info
func (u *Event) Update(db *sql.DB, user *eventers.Event) error {
	//return db.Update(user)
	return nil
}

// List returns list of all events retrievable for the current event, depending on role
func (u *Event) List(db *sql.DB, p *eventers.Pagination) ([]eventers.Event, error) {
	return []eventers.Event{}, nil
}

// Delete an event
func (u *Event) Delete(db *sql.DB, user *eventers.Event) error {
	return nil
}
