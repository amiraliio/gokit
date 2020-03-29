package main

import (
	"context"
	"log"
	"net/http"
	"os"

	"github.com/amiraliio/gokit/config"
	"github.com/amiraliio/gokit/todo"
)

func main() {

	root, err := os.Getwd()
	if err != nil {
		log.Fatalln(err)
	}

	sys := config.InitConfig(root)

	env := sys.Config()

	logger := sys.Logger()

	db := sys.DB()
	defer db.Close()

	if env.GetBool("APP.DEBUG.ENABLED") {
		sys.Profiler(env.GetString("APP.DEBUG.PORt"))
	}

	repository := todo.NewMysqlRepository(db)

	var service todo.Service

	service = todo.NewService(repository)

	service = todo.NewLoggerService(logger, service)

	endpoint := todo.NewEndpoint(service)

	transport := todo.NewTransport(context.Background(), endpoint)

	if err := http.ListenAndServe(":"+env.GetString("APP.PORT"), transport); err != nil {
		log.Fatalln(err)
	}

}
