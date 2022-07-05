package account

import (
	"context"
	"github.com/go-kit/log"
	"github.com/go-kit/log/level"

	"github.com/gofrs/uuid"
)

type service struct {
	repo   Repository
	logger log.Logger
}

func NewService(logger log.Logger, repo Repository) Service {
	return &service{repo, logger}
}

func (s *service) CreateUser(c context.Context, email string, name string) (string, error) {
	logger := log.With(s.logger, "logic", "CreateUser")

	id, _ := uuid.NewV4()
	var user User
	user.ID = id.String()
	user.Name = name
	user.Email = email

	if err := s.repo.CreateUser(c, user); err != nil {
		level.Error(logger).Log("err", err)
		return "", err
	}

	logger.Log("CreatedUser", id)

	return "Success", nil
}

func (s *service) GetUser(c context.Context, id string) (string, error) {
	logger := log.With(s.logger, "logic", "GetUser")
	email, err := s.repo.GetUser(c, id)
	if err != nil {
		level.Error(logger).Log("err", err)
		return "", err
	}

	logger.Log("GetUser", id)

	return email, nil
}
