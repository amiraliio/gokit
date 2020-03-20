package todo

import (
	"context"
	"database/sql"
	"fmt"
	"time"

	"github.com/go-kit/kit/log"
	"github.com/go-kit/kit/log/level"
	"github.com/google/uuid"
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
	cursor, err := r.db.Query("select title, text from todo")
	if err != nil {
		level.Error(r.logger).Log("Repository", "Todo", "PostgreSQL", "List", err.Error())
		return nil, err
	}
	defer cursor.Close()
	var list []*TODO
	for cursor.Next() {
		var todo *TODO
		if err := cursor.Scan(&todo.Title, &todo.Text); err != nil {
			fmt.Println(err)
			level.Error(r.logger).Log("Repository", "Todo", "PostgreSQL", "List", err.Error())
			return nil, err
		}
		list = append(list, todo)
	}
	r.logger.Log("Repository", "Todo", "PostgreSQL", "List", "Success")
	return list, nil
}

func (r *repository) Insert(ctx context.Context, title, text string) error {
	_, err := r.db.Exec("insert into todo(id, title, text, created_at) values($1, $2, $3, $4)", uuid.New(), title, text, time.Now())
	if err != nil {
		level.Error(r.logger).Log("Repository", "Todo", "PostgreSQL", "Insert", err.Error())
		return err
	}
	r.logger.Log("Repository", "Todo", "PostgreSQL", "Insert", "Success")
	return nil
}
