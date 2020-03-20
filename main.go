package main

import (
	"context"
	"database/sql"
	"net/http"
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

	transport := todo.NewTransport(context.Background(), endpoint)

	if err := http.ListenAndServe(":8976", transport); err != nil {
		level.Error(logger).Log(err.Error)
	}

}