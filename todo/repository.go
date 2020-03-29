package todo

import (
	"database/sql"
	"github.com/amiraliio/gokit/helper"
	"time"
)

type repository struct {
	db *sql.DB
}

func NewMysqlRepository(db *sql.DB) Repository {
	return &repository{
		db: db,
	}
}

func (r *repository) List() ([]*TODO, error) {
	cursor, err := r.db.Query("select id, title, text, created_at from todo")
	if err != nil {
		return nil, err
	}
	defer cursor.Close()
	var list []*TODO
	for cursor.Next() {
		todo := new(TODO)
		if err := cursor.Scan(&todo.ID, &todo.Title, &todo.Text, &todo.Create_at); err != nil {
			return nil, err
		}
		list = append(list, todo)
	}
	return list, nil
}

func (r *repository) Insert(title, text string) error {
	_, err := r.db.Exec("insert into todo(id, title, text, created_at) values($1, $2, $3, $4)", helper.NewUUID(), title, text, time.Now())
	if err != nil {
		return err
	}
	return nil
}
