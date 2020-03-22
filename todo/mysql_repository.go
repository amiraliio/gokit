package todo

import (
	"context"
	"database/sql"
	"time"

	"github.com/go-kit/kit/log"
	"github.com/go-kit/kit/log/level"
	"github.com/google/uuid"
)


type repository struct {
	db     *sql.DB
	logger log.Logger
}

func NewMysqlRepository(db *sql.DB, logger log.Logger) Repository {
	return &repository{
		db:     db,
		logger: log.With(logger, "Repository", "Todo", "PostgreSQL"),
	}
}

func (r *repository) List(ctx context.Context) ([]*TODO, error) {
	defer r.db.Close()
	cursor, err := r.db.Query("select id, title, text, created_at from todo")
	if err != nil {
		level.Error(r.logger).Log("Repository", "Todo", "PostgreSQL", "List", err.Error())
		return nil, err
	}
	defer cursor.Close()
	var list []*TODO
	for cursor.Next() {
		todo := new(TODO)
		if err := cursor.Scan(&todo.ID, &todo.Title, &todo.Text, &todo.Create_at); err != nil {
			level.Error(r.logger).Log("Repository", "Todo", "PostgreSQL", "List", err.Error())
			return nil, err
		}
		list = append(list, todo)
	}
	r.logger.Log("Repository", "Todo", "PostgreSQL", "List", "Success")
	return list, nil
}

func (r *repository) Insert(ctx context.Context, title, text string) error {
	defer r.db.Close()
	_, err := r.db.Exec("insert into todo(id, title, text, created_at) values($1, $2, $3, $4)", uuid.New(), title, text, time.Now())
	if err != nil {
		level.Error(r.logger).Log("Repository", "Todo", "PostgreSQL", "Insert", err.Error())
		return err
	}
	r.logger.Log("Repository", "Todo", "PostgreSQL", "Insert", "Success")
	return nil
}
