package account

import (
	"context"
	"database/sql"

	"github.com/go-kit/kit/log"
)

type repository struct {
	db     *sql.DB
	logger log.Logger
}

func NewRepository(db *sql.DB, logger log.Logger) Repository {
	return &repository{
		db,
		log.With(logger, "Repository", "SQL"),
	}
}

func (r *repository) CreateUser(ctx context.Context, user *User) error {
	//TODO implement insert into DB
	return nil
}

func (r *repository) GetUser(ctx context.Context, id string) (string, error) {
	var email string
	if err := r.db.QueryRow("select `email` from `users` where `id`=?", id).Scan(&email); err != nil {
		return "", err
	}
	return email, nil
}
