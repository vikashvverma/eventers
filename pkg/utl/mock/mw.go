package mock

import (
	"github.com/vikashvverma/eventers/pkg/utl/model"
)

// JWT mock
type JWT struct {
	GenerateTokenFn func(*eventers.User) (string, string, error)
}

// GenerateToken mock
func (j *JWT) GenerateToken(u *eventers.User) (string, string, error) {
	return j.GenerateTokenFn(u)
}
