package app

import (
	"database/sql"
	"fmt"
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
	log "github.com/sirupsen/logrus"
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

	// Logger
	SetLogrus(cfg.Log.Level)

	// Postgres
	db, err := NewPostgres(cfg.PG.URL)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	// Repository
	rep := repository.NewRepositories(db)

	//Service dependencies
	deps := service.NewServicesDeps(rep, hasher.NewSHA1Hasher(cfg.Salt))

	// Services
	services := service.NewServices(deps)

	// mux handler
	log.Info("configuring router...")
	handler := mux.NewRouter()
	v1.New(handler, services)

	// HTTP Server
	log.Info("starting server...")
	httpserver := httpserver.New(handler, cfg.Port)

	// Waiting signal
	log.Info("configuring graceful shutdown...")
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
