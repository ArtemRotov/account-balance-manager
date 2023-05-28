package app

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/ArtemRotov/account-balance-manager/config"
	v1 "github.com/ArtemRotov/account-balance-manager/internal/controller/http/v1"
	"github.com/ArtemRotov/account-balance-manager/pkg/httpserver"
	"github.com/gorilla/mux"
	log "github.com/sirupsen/logrus"
)

func Run(configPath string) {
	cfg, err := config.NewConfig(configPath)
	if err != nil {
		fmt.Errorf("Config error: %s", err)
	}
	log.SetLevel(log.DebugLevel)
	handler := mux.NewRouter()
	v1.New(handler)

	httpserver := httpserver.New(handler, cfg.Port)

	// Waiting signal
	log.Info("Configuring graceful shutdown...")
	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, os.Interrupt, syscall.SIGTERM)

	select {
	case s := <-interrupt:
		log.Info("app - Run - signal: " + s.String())
	case err = <-httpserver.Notify():
		log.Error(fmt.Errorf("app - Run - httpServer.Notify: %w", err))
	}

	// Graceful shutdown
	log.Info("Shutting down...")
	err = httpserver.Shutdown()
	if err != nil {
		log.Error(fmt.Errorf("app - Run - httpServer.Shutdown: %w", err))
	}
}
