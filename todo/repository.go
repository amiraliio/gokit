package todo

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
		db:     db,
		logger: log.With(logger, "Repository", "Todo", "PostgreSQL"),
	}
}

func (r *repository) List(ctx context.Context) ([]*TODO, error) {
	return nil, nil
}

func (r *repository) Insert(ctx context.Context, title, text string) error {
	return nil
}
