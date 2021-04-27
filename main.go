package main

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/quangdangfit/go-backend/app"
	"github.com/quangdangfit/go-backend/config"
	"github.com/quangdangfit/go-backend/pkg/logger"
)

func main() {
	conf := config.GetConfig()
	logger.Initialize(conf.Env)

	container := app.BuildContainer()
	engine := app.InitGinEngine(container)

	server := &http.Server{
		Addr:    fmt.Sprintf(":%d", conf.Port),
		Handler: engine,
	}

	go func() {
		// service connections
		logger.Info("Listen at:", server.Addr)
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			logger.Fatalf("Error: %s\n", err)
		}
	}()

	// Wait for interrupt signal to gracefully shutdown the server with
	// a timeout of 5 seconds.
	quit := make(chan os.Signal)
	// kill (no param) default send syscanll.SIGTERM
	// kill -2 is syscall.SIGINT
	// kill -9 is syscall. SIGKILL but can"t be catch, so don't need add it
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	logger.Info("Shutdown Server ...")

	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
	defer cancel()
	if err := server.Shutdown(ctx); err != nil {
		logger.Fatal("Server Shutdown: ", err)
	}
	// catching ctx.Done(). timeout of 1 seconds.
	select {
	case <-ctx.Done():
		logger.Info("Timeout of 1 seconds.")
	}
	logger.Info("Server exiting")
}
