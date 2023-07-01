package app

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/ArtemRotov/account-balance-manager/config"
	v1 "github.com/ArtemRotov/account-balance-manager/internal/controller/http/v1"
	"github.com/ArtemRotov/account-balance-manager/internal/repository"
	"github.com/ArtemRotov/account-balance-manager/internal/service"
	"github.com/ArtemRotov/account-balance-manager/pkg/hasher"
	"github.com/ArtemRotov/account-balance-manager/pkg/httpserver"
	"github.com/gorilla/mux"
	_ "github.com/lib/pq"
)

// @title My test service
// @version 1.0
// @description My test service
// @termsOfService http://swagger.io/terms/

// @contact.name Artem Rotov
// @contact.url https://github.com/ArtemRotov
// @contact.email rotoffff@yandex.ru

// @host 127.0.0.1:8080
// @BasePath /

func Run(configPath string) {

	cfg, err := config.NewConfig(configPath)
	if err != nil {
		log.Fatal(err)
	}

	// Postgres
	db, err := NewPostgres(cfg.PG.URL)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	// Logger
	logger := SetSlog(cfg.Level)

	// Repository
	rep := repository.NewRepositories(db)

	//Service dependencies
	deps := service.NewServicesDeps(rep, logger, hasher.NewSHA1Hasher(cfg.Salt), cfg.SignKey, cfg.TokenTTL)

	// Services
	services := service.NewServices(deps)

	// mux handler
	logger.Info("configuring router...")
	handler := mux.NewRouter()
	v1.New(handler, services, logger)

	// HTTP Server
	logger.Info("starting server...")
	server := httpserver.New(handler, cfg.Port)

	// Waiting signal
	logger.Info("configuring graceful shutdown...")
	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, os.Interrupt, syscall.SIGTERM)

	select {
	case s := <-interrupt:
		logger.Info("app - Run - signal: " + s.String())
	case err = <-server.Notify():
		logger.Error(fmt.Sprintf("app - Run - httpServer.Notify: %w", err))
	}

	// Graceful shutdown
	logger.Info("Shutting down...")
	err = server.Shutdown()
	if err != nil {
		logger.Error(fmt.Sprintf("app - Run - httpServer.Shutdown: %w", err))
	}
}

func NewPostgres(url string) (*sql.DB, error) {
	db, err := sql.Open("postgres", url)
	if err != nil {
		return nil, err
	}

	if err := db.Ping(); err != nil {
		return nil, err
	}

	return db, nil
}
