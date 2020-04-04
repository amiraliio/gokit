package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/amiraliio/gokit/config"
	"github.com/amiraliio/gokit/todo"
	kitprometheus "github.com/go-kit/kit/metrics/prometheus"
	stdprometheus "github.com/prometheus/client_golang/prometheus"
)

func main() {
// rate limiting
	// metrics
	// load balancing
	// analytics
	// logging
	// circuit breaking
	// grpc client and server
	// service mesh

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

	service = todo.NewMetricsService(kitprometheus.NewCounterFrom(
		stdprometheus.CounterOpts{
			Namespace: "api",
			Subsystem: "todo_service",
			Name:      "request_count",
			Help:      "Number of requests received.",
		}, []string{"method"}),
		kitprometheus.NewSummaryFrom(stdprometheus.SummaryOpts{
			Namespace: "api",
			Subsystem: "todo_service",
			Name:      "request_latency_microseconds",
			Help:      "Total duration of requests in microseconds.",
		}, []string{"method"}),
		service)

	endpoint := todo.NewEndpoint(service)

	transport := todo.NewTransport(context.Background(), endpoint)

	errChannel := make(chan error, 2)

	go func(handler http.Handler) {
		fmt.Println("Listening on port " + env.GetString("APP.PORT") + "...")
		errChannel <- http.ListenAndServe(":"+env.GetString("APP.PORT"), handler)
	}(transport)

	go func() {
		c := make(chan os.Signal, 1)
		signal.Notify(c, syscall.SIGINT)
		errChannel <- fmt.Errorf("%s", <-c)
	}()

	fmt.Println(<-errChannel)
}
