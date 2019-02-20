package user

import (
	"github.com/go-pg/pg"
	"github.com/go-pg/pg/orm"
	"github.com/labstack/echo"
	"github.com/vikashvverma/eventers/pkg/api/user/platform/pgsql"
	"github.com/vikashvverma/eventers/pkg/utl/model"
)

// Service represents user application interface
type Service interface {
	Create(echo.Context, eventers.User) (*eventers.User, error)
	List(echo.Context, *eventers.Pagination) ([]eventers.User, error)
	View(echo.Context, int) (*eventers.User, error)
	Delete(echo.Context, int) error
	Update(echo.Context, *Update) (*eventers.User, error)
}

// New creates new user application service
func New(db *pg.DB, udb UDB, rbac RBAC, sec Securer) *User {
	return &User{db: db, udb: udb, rbac: rbac, sec: sec}
}

// Initialize initalizes User application service with defaults
func Initialize(db *pg.DB, rbac RBAC, sec Securer) *User {
	return New(db, pgsql.NewUser(), rbac, sec)
}

// User represents user application service
type User struct {
	db   *pg.DB
	udb  UDB
	rbac RBAC
	sec  Securer
}

// Securer represents security interface
type Securer interface {
	Hash(string) string
}

// UDB represents user repository interface
type UDB interface {
	Create(orm.DB, eventers.User) (*eventers.User, error)
	View(orm.DB, int) (*eventers.User, error)
	List(orm.DB, *eventers.ListQuery, *eventers.Pagination) ([]eventers.User, error)
	Update(orm.DB, *eventers.User) error
	Delete(orm.DB, *eventers.User) error
}

// RBAC represents role-based-access-control interface
type RBAC interface {
	User(echo.Context) *eventers.AuthUser
	EnforceUser(echo.Context, int) error
	AccountCreate(echo.Context, eventers.AccessRole, int, int) error
	IsLowerRole(echo.Context, eventers.AccessRole) error
}
