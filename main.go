package main

import (
	"context"
	"database/sql"
	"os"

	"github.com/amiraliio/gokit/todo"
	"github.com/go-kit/kit/log"
	"github.com/go-kit/kit/log/level"
	_ "github.com/lib/pq"
)

func main() {
	logger := log.NewLogfmtLogger(log.NewSyncWriter(os.Stderr))

	db, err := sql.Open("postgres", "user=postgres dbname=gokitTodo sslmode=verify-full")
	if err != nil {
		level.Error(logger).Log("DB", err.Error())
		return
	}
	repository := todo.NewRepository(db, logger)

	service := todo.NewService(repository, logger)

	endpoint := todo.NewEndpoint(service)

	todo.NewTransport(context.Background(), endpoint)

}
