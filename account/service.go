package account

import "context"

// Service 层
type Service interface {
	CreateUser(c context.Context, email string, name string) (string, error)
	GetUser(c context.Context, id string) (string, error)
}
