package event

import (
	"github.com/labstack/echo"

	"github.com/vikashvverma/eventers/pkg/utl/model"
)

// Create creates a new event
func (u *Event) Create(c echo.Context, req eventers.Event) (*eventers.Event, error) {
	return u.udb.Create(u.db, req)
}

// View returns single event
func (u *Event) View(c echo.Context, id int) (*eventers.Event, error) {
	return u.udb.View(u.db, id)
}
