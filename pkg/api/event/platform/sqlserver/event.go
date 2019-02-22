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

// Event represents the client for event table
type Event struct{}

// Create creates a new event on database
func (u *Event) Create(db *sql.DB, e eventers.Event) (*eventers.Event, error) {
	ctx := context.Background()
	var err error

	// Check if database is alive.
	err = db.PingContext(ctx)
	if err != nil {
		return nil, err
	}

	tsql := `INSERT INTO Event (Name, Date, Location) VALUES (@Name, @Date, @Location);
			 SELECT convert(bigint, SCOPE_IDENTITY());`

	// Execute query
	stmt, err := db.Prepare(tsql)
	if err != nil {
		return nil, err
	}
	defer stmt.Close()

	row := stmt.QueryRowContext(
		ctx,
		sql.Named("Name", e.Name),
		sql.Named("Date", e.Date),
		sql.Named("Location", e.Location),
	)

	if row == nil {
		return nil, fmt.Errorf("invalid data")
	}
	var newID int64
	err = row.Scan(&newID)
	if err != nil {
		return nil, fmt.Errorf("unable to create event, invalid data provided")
	}

	return &eventers.Event{ID: newID}, nil
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
