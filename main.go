package main

import (
	"context"
	"database/sql"
	"flag"
	"fmt"
	"github.com/go-kit/log"
	"github.com/go-kit/log/level"
	_ "github.com/go-sql-driver/mysql"
	"gokit-account/account"
	"net/http"
	"os"
	"os/signal"
	"syscall"
)

var dbsource = "root:root@tcp(127.0.0.1:3306)/recordings?charset=utf8mb4&parseTime=True&loc=Local"

func main() {
	var httpAddr = flag.String("http", ":8081", "http listen address")
	// go-kit logger
	var logger log.Logger
	{
		logger = log.NewLogfmtLogger(os.Stderr)
		logger = log.NewSyncLogger(logger)
		logger = log.With(logger,
			"service", "account",
			"time:", log.DefaultTimestampUTC,
			"caller", log.DefaultCaller,
		)
	}
	level.Info(logger).Log("msg", "service started")
	defer level.Info(logger).Log("msg", "service ended")

	// database
	var db *sql.DB
	{
		var err error
		db, err = sql.Open("mysql", dbsource)
		if err != nil {
			level.Error(logger).Log("exit", err)
			os.Exit(-1)
		}
	}

	// Service
	flag.Parse()
	ctx := context.Background()
	var srv account.Service
	{
		repository := account.NewRepo(db, logger)
		srv = account.NewService(logger, repository)
	}

	// error interpret
	errs := make(chan error)
	go func() {
		c := make(chan os.Signal, 1)
		signal.Notify(c, syscall.SIGINT, syscall.SIGTERM)
		errs <- fmt.Errorf("%s", <-c)
	}()

	// endpoints
	endpoints := account.MakeEndpoints(srv)
	go func() {
		fmt.Println("listening on port", *httpAddr)
		handler := account.NewHttpServer(ctx, endpoints)
		errs <- http.ListenAndServe(*httpAddr, handler)
	}()

	level.Error(logger).Log("exit", <-errs)
}
