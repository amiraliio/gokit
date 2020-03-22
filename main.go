package main

import (
	"context"
	"log"
	"net/http"
	"os"

	"github.com/amiraliio/gokit/config"
	"github.com/amiraliio/gokit/todo"
	_ "net/http/pprof"
)

func main() {

	root, err := os.Getwd()
	if err != nil {
		log.Fatalln(err)
	}

	sys := config.InitConfig(root)

	env := sys.Config()

	logger := sys.Logger()

	if env.GetBool("APP.DEBUG.ENABLED") {
		sys.Profiler(env.GetString("APP.DEBUG.PORt"))
	}

	repository := todo.NewMysqlRepository(sys.DB(), logger)

	service := todo.NewService(repository, logger)

	endpoint := todo.NewEndpoint(service)

	transport := todo.NewTransport(context.Background(), endpoint)

	if err := http.ListenAndServe(":"+env.GetString("APP.PORT"), transport); err != nil {
		log.Fatalln(err)
	}

}
