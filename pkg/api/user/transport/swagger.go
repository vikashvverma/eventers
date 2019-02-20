package transport

import (
	"github.com/vikashvverma/eventers/pkg/utl/model"
)

// User model response
// swagger:response userResp
type swaggUserResponse struct {
	// in:body
	Body struct {
		*eventers.User
	}
}

// Users model response
// swagger:response userListResp
type swaggUserListResponse struct {
	// in:body
	Body struct {
		Users []eventers.User `json:"users"`
		Page  int          `json:"page"`
	}
}
