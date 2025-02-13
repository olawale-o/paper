package main

import (
	"context"
	"flag"
	"fmt"
	"go-simple-rest/db"
	"go-simple-rest/src/v1/auth"
	"go-simple-rest/src/v1/auth/impl"
	"go-simple-rest/src/v1/auth/mongodb"
	"go-simple-rest/src/v1/auth/transport"
	httptransport "go-simple-rest/src/v1/auth/transport/http"
	"net/http"
	"os"
	"os/signal"
	"sync"
	"syscall"

	"github.com/gin-gonic/gin"
	"github.com/go-kit/log"
	"github.com/go-kit/log/level"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

func terminate() {
	if r := recover(); r != nil {
		fmt.Println("An error occured ", r)
		fmt.Println("Application terminated gracefully")
	} else {
		fmt.Println("Application executed succcesfully")
	}
}

func iterateChangeStream(routineCtx context.Context, waitGroup sync.WaitGroup, stream *mongo.ChangeStream) {
	defer stream.Close(routineCtx)
	defer waitGroup.Done()
	for stream.Next(routineCtx) {
		var data bson.M
		if err := stream.Decode(&data); err != nil {
			panic(err)
		}
		fmt.Printf("%v\n", data)
	}
}

var client, ctx, err = db.Connect()

func main() {
	var (
		httpAddr = flag.String("http.addr", ":8080", "HTTP listen address")
	)
	var logger log.Logger
	{
		logger = log.NewLogfmtLogger(os.Stderr)
		logger = log.NewSyncLogger(logger)
		logger = level.NewFilter(logger, level.AllowDebug())
		logger = log.With(logger,
			"svc", "order",
			"ts", log.DefaultTimestampUTC,
			"caller", log.DefaultCaller,
		)
	}
	var svc auth.Service
	{

		repo, err := mongodb.New(client.Database("go"), logger)
		if err != nil {
			level.Error(logger).Log("exit", err)
			os.Exit(-1)
		}
		svc = impl.NewService(repo, logger)
	}

	var h *gin.Engine
	{
		endpoints := transport.MakeEndpoints(svc)
		h = httptransport.NewService(endpoints, logger)
	}

	errs := make(chan error)
	go func() {
		c := make(chan os.Signal)
		signal.Notify(c, syscall.SIGINT, syscall.SIGTERM)
		errs <- fmt.Errorf("%s", <-c)
	}()

	go func() {
		level.Info(logger).Log("transport", "HTTP", "addr", *httpAddr)
		server := &http.Server{
			Addr:    *httpAddr,
			Handler: h,
		}
		errs <- server.ListenAndServe()
	}()

	level.Error(logger).Log("exit", <-errs)

}
