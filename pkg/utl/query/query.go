package query

import (
	"github.com/labstack/echo"
	"github.com/vikashvverma/eventers/pkg/utl/model"
)

// List prepares data for list queries
func List(u *eventers.AuthUser) (*eventers.ListQuery, error) {
	switch true {
	case u.Role <= eventers.AdminRole: // user is SuperAdmin or Admin
		return nil, nil
	case u.Role == eventers.CompanyAdminRole:
		return &eventers.ListQuery{Query: "company_id = ?", ID: u.CompanyID}, nil
	case u.Role == eventers.LocationAdminRole:
		return &eventers.ListQuery{Query: "location_id = ?", ID: u.LocationID}, nil
	default:
		return nil, echo.ErrForbidden
	}
}
