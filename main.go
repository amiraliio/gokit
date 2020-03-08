package main

import (
	"context"
	"database/sql"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/go-kit/kit/log"

	"github.com/amiraliio/gokit/account"
)

func main() {

	logger := log.NewLogfmtLogger(log.NewSyncWriter(os.Stderr))

	repository := account.NewRepository(&sql.DB{}, logger)

	service := account.NewService(repository, logger)

	endpoints := account.MakeEndpoints(service)

	account.NewHTTPServer(context.Background(), endpoints)

	err := make(chan error, 2)

	go func() {
		err <- http.ListenAndServe(":8765", nil)
	}()

	go func() {
		c := make(chan os.Signal)
		signal.Notify(c, syscall.SIGINT)
		err <- fmt.Errorf("%s", <-c)
	}()

	logger.Log(<-err)
}
