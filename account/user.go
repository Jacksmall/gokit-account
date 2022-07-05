package account

import "context"

// User struct
type User struct {
	ID    string `json:"id,omitempty"`
	Email string `json:"email"`
	Name  string `json:"name"`
}

// Repository interface
type Repository interface {
	CreateUser(c context.Context, user User) error
	GetUser(c context.Context, id string) (string, error)
}
