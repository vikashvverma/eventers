package mockdb

import (
	"github.com/go-pg/pg/orm"
	"github.com/vikashvverma/eventers/pkg/utl/model"
)

// User database mock
type User struct {
	CreateFn         func(orm.DB, eventers.User) (*eventers.User, error)
	ViewFn           func(orm.DB, int) (*eventers.User, error)
	FindByUsernameFn func(orm.DB, string) (*eventers.User, error)
	FindByTokenFn    func(orm.DB, string) (*eventers.User, error)
	ListFn           func(orm.DB, *eventers.ListQuery, *eventers.Pagination) ([]eventers.User, error)
	DeleteFn         func(orm.DB, *eventers.User) error
	UpdateFn         func(orm.DB, *eventers.User) error
}

// Create mock
func (u *User) Create(db orm.DB, usr eventers.User) (*eventers.User, error) {
	return u.CreateFn(db, usr)
}

// View mock
func (u *User) View(db orm.DB, id int) (*eventers.User, error) {
	return u.ViewFn(db, id)
}

// FindByUsername mock
func (u *User) FindByUsername(db orm.DB, uname string) (*eventers.User, error) {
	return u.FindByUsernameFn(db, uname)
}

// FindByToken mock
func (u *User) FindByToken(db orm.DB, token string) (*eventers.User, error) {
	return u.FindByTokenFn(db, token)
}

// List mock
func (u *User) List(db orm.DB, lq *eventers.ListQuery, p *eventers.Pagination) ([]eventers.User, error) {
	return u.ListFn(db, lq, p)
}

// Delete mock
func (u *User) Delete(db orm.DB, usr *eventers.User) error {
	return u.DeleteFn(db, usr)
}

// Update mock
func (u *User) Update(db orm.DB, usr *eventers.User) error {
	return u.UpdateFn(db, usr)
}
