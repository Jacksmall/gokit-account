package account

import (
	"context"
	"database/sql"
	"github.com/go-kit/log"
	"github.com/go-kit/log/level"
)

type repo struct {
	db     *sql.DB
	logger log.Logger
}

func NewRepo(db *sql.DB, logger log.Logger) Repository {
	return &repo{db, logger}
}

func (r *repo) CreateUser(c context.Context, u User) error {
	logger := log.With(r.logger, "repo", "CreateUser")

	sql := `INSERT INTO account (id, email, name) VALUES (?, ?, ?)`
	_, err := r.db.ExecContext(c, sql, u.ID, u.Email, u.Name)
	if err != nil {
		level.Error(logger).Log("repo", "CreateUser", err)
		return err
	}
	logger.Log("repo", "CreateUser", "Success")

	return nil
}

func (r *repo) GetUser(c context.Context, id string) (string, error) {
	logger := log.With(r.logger, "repo", "GetUser")

	var email string
	if err := r.db.QueryRowContext(c, "SELECT email FROM account WHERE id=?", id).Scan(&email); err != nil {
		level.Error(logger).Log("repo", "GetUser", err)
		return "", err
	}
	logger.Log("repo", "GetUser", "Success")

	return email, nil
}
