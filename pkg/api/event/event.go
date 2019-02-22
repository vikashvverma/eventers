package event

import (
	"github.com/denisenkom/go-mssqldb"
	"github.com/labstack/echo"

	"github.com/vikashvverma/eventers/pkg/utl/model"
	"github.com/vikashvverma/eventers/pkg/utl/structs"
)

// Create creates a new user account
func (u *Event) Create(c echo.Context, req eventers.Event) (int64, error) {
	return u.udb.Create(u.db, req)
}

// List returns list of users
func (u *Event) List(c echo.Context, p *eventers.Pagination) ([]eventers.Event, error) {
	return u.udb.List(u.db, p)
}

// View returns single user
func (u *Event) View(c echo.Context, id int) (*eventers.Event, error) {
	return u.udb.View(u.db, id)
}

// Delete deletes a user
func (u *Event) Delete(c echo.Context, id int) error {
	user, err := u.udb.View(u.db, id)
	if err != nil {
		return err
	}
	return u.udb.Delete(u.db, user)
}

// Update contains event's information used for updating
type Update struct {
	ID       int
	Name     *string
	Date     *mssql.DateTime1
	Location *string
}

// Update updates user's contact information
func (u *Event) Update(c echo.Context, req *Update) (*eventers.Event, error) {
	user, err := u.udb.View(u.db, req.ID)
	if err != nil {
		return nil, err
	}

	structs.Merge(user, req)
	if err := u.udb.Update(u.db, user); err != nil {
		return nil, err
	}

	return user, nil
}
