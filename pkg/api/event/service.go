package event

import (
	"database/sql"

	"github.com/labstack/echo"

	"github.com/vikashvverma/eventers/pkg/api/event/platform/sqlserver"
	"github.com/vikashvverma/eventers/pkg/utl/model"
)

// Service represents event application interface
type Service interface {
	Create(echo.Context, eventers.Event) (int64, error)
	List(echo.Context, *eventers.Pagination) ([]eventers.Event, error)
	View(echo.Context, int) (*eventers.Event, error)
	Delete(echo.Context, int) error
	Update(echo.Context, *Update) (*eventers.Event, error)
}

// New creates new event application service
func New(db *sql.DB, udb UDB, sec Securer) *Event {
	return &Event{db: db, udb: udb, sec: sec}
}

// Initialize initalizes Event application service with defaults
func Initialize(db *sql.DB, sec Securer) *Event {
	return New(db, sqlserver.NewEvent(), sec)
}

// User represents event application service
type Event struct {
	db  *sql.DB
	udb UDB
	sec Securer
}

// Securer represents security interface
type Securer interface {
	Hash(string) string
}

// UDB represents event repository interface
type UDB interface {
	Create(*sql.DB, eventers.Event) (int64, error)
	View(*sql.DB, int) (*eventers.Event, error)
	List(*sql.DB, *eventers.Pagination) ([]eventers.Event, error)
	Update(*sql.DB, *eventers.Event) error
	Delete(*sql.DB, *eventers.Event) error
}
