package config

import (
	"database/sql"
	syslog "log"
	"net/http"
	"net/http/pprof"
	"os"

	"github.com/go-kit/kit/log"
	"github.com/gorilla/mux"
	_ "github.com/lib/pq"
	"github.com/spf13/viper"
)

type Config interface {
	Config() *viper.Viper
	DB() *sql.DB
	Logger() log.Logger
	Profiler(debugPort string)
}

func InitConfig(root string) Config {
	return &config{
		root: root,
	}
}

type config struct {
	root string
}

func (c *config) Config() *viper.Viper {
	initConfig := viper.New()
	initConfig.AddConfigPath(c.root)
	if err := initConfig.ReadInConfig(); err != nil {
		syslog.Fatalln("Config", "Viper", err.Error())
	}
	return initConfig
}

func (c *config) DB() *sql.DB {
	db, err := sql.Open("postgres", "user=postgres dbname=gokitTodo sslmode=disable")
	if err != nil {
		syslog.Fatalln("Config", "DB", err.Error())
	}
	return db
}

func (c *config) Logger() log.Logger {
	return log.NewLogfmtLogger(log.NewSyncWriter(os.Stderr))
}

func (c *config) Profiler(debugPort string) {
	go func() {
		debugR := mux.NewRouter()
		debugR.HandleFunc("/debug/pprof", pprof.Index)
		debugR.HandleFunc("/debug/pprof/cmdline", pprof.Cmdline)
		debugR.HandleFunc("/debug/pprof/profile", pprof.Profile) //cpu profile
		debugR.HandleFunc("/debug/pprof/symbol", pprof.Symbol)
		debugR.HandleFunc("/debug/pprof/trace", pprof.Trace)        //execution trace - go tool trace
		debugR.Handle("/debug/pprof/mutex", pprof.Handler("mutex")) //holders of contended mutexes
		debugR.Handle("/debug/pprof/heap", pprof.Handler("heap"))   //a sampling of all heap allocations
		debugR.Handle("/debug/pprof/goroutine", pprof.Handler("goroutine")) //stack traces of all current goroutines
		debugR.Handle("/debug/pprof/block", pprof.Handler("block")) //stack traces that led to blocking on synchronization primitives
		debugR.Handle("/debug/pprof/threadcreate", pprof.Handler("threadcreate")) //stack traces that led to the creation of new OS threads
		if err := http.ListenAndServe(":"+debugPort, debugR); err != nil {
			syslog.Fatalln(err)
		}
	}()
}
