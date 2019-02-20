package query_test

import (
	"testing"

	"github.com/vikashvverma/eventers/pkg/utl/model"

	"github.com/labstack/echo"

	"github.com/vikashvverma/eventers/pkg/utl/query"
	"github.com/stretchr/testify/assert"
)

func TestList(t *testing.T) {
	type args struct {
		user *eventers.AuthUser
	}
	cases := []struct {
		name     string
		args     args
		wantData *eventers.ListQuery
		wantErr  error
	}{
		{
			name: "Super admin user",
			args: args{user: &eventers.AuthUser{
				Role: eventers.SuperAdminRole,
			}},
		},
		{
			name: "Company admin user",
			args: args{user: &eventers.AuthUser{
				Role:      eventers.CompanyAdminRole,
				CompanyID: 1,
			}},
			wantData: &eventers.ListQuery{
				Query: "company_id = ?",
				ID:    1},
		},
		{
			name: "Location admin user",
			args: args{user: &eventers.AuthUser{
				Role:       eventers.LocationAdminRole,
				CompanyID:  1,
				LocationID: 2,
			}},
			wantData: &eventers.ListQuery{
				Query: "location_id = ?",
				ID:    2},
		},
		{
			name: "Normal user",
			args: args{user: &eventers.AuthUser{
				Role: eventers.UserRole,
			}},
			wantErr: echo.ErrForbidden,
		},
	}
	for _, tt := range cases {
		t.Run(tt.name, func(t *testing.T) {
			q, err := query.List(tt.args.user)
			assert.Equal(t, tt.wantData, q)
			assert.Equal(t, tt.wantErr, err)
		})
	}
}
