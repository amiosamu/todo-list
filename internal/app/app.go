package app

import (
	"context"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/amiosamu/todo-list/config"
	todo_list "github.com/amiosamu/todo-list/internal/controller/http/todo-list"
	"github.com/amiosamu/todo-list/internal/repo"
	"github.com/amiosamu/todo-list/internal/service"
	"github.com/amiosamu/todo-list/pkg/httpserver"
	"github.com/amiosamu/todo-list/pkg/mongo"
	"github.com/gin-gonic/gin"
	mongo2 "go.mongodb.org/mongo-driver/mongo"
	_ "github.com/amiosamu/todo-list/docs"
	"golang.org/x/exp/slog"
)


// @title To-Do List API. 
// @version 1.0
// @description Тестовое задание на позицию Junior Go разработчика в компанию ТОО Region LLC.
// @host localhost:8080
// @BasePath /

func Run(path string) {
	cfg, err := config.NewConfig(path)
	if err != nil {
		slog.Error("error reading config: %w", err)
	}
	slog.Info("setting up MongoDB...")
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*1)
	defer cancel()
	mg, err := mongo.InitMongo(ctx, cfg)
	if err != nil {
		log.Fatalf("failed connecting to Mongo: %v", err)
	}
	ctxDB, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	if err := mg.Ping(ctxDB, nil); err != nil {
		log.Fatalf("error pinging Mongo: %v", err)
	}
	defer func(ctx context.Context, client *mongo2.Client) {
		err := mongo.ShutdownDB(ctx, client)
		if err != nil {
			log.Fatalf("error shuttding down Mongo: %v", err)
		}
	}(ctxDB, mg)

	slog.Info("initializing repositories...")

	repository := repo.NewRepos(mg)

	slog.Info("initializing service dependencies...")

	dependencies := service.Dependencies{

		Repos: repository,
	}

	services := service.NewServices(dependencies)

	slog.Info("initializing handlers and routes...")

	handler := gin.New()
	todo_list.NewRouter(handler, services)

	slog.Info("starting http server...")
	slog.Debug("server port: %s", cfg.HTTP.Port)
	httpServer := httpserver.New(handler, httpserver.Port(cfg.HTTP.Port))

	slog.Info("configuring graceful shutdown...")
	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, os.Interrupt, syscall.SIGTERM)

	select {
	case s := <-interrupt:
		slog.Info("app - Run - signal: " + s.String())
	case err = <-httpServer.Notify():
		slog.Error("app - Run - httpServer.Notify: " + err.Error())
	}

	slog.Info("shutting down...")
	quit := make(chan os.Signal)
	signal.Notify(quit, syscall.SIGTERM, syscall.SIGINT)
	err = httpServer.Shutdown()
	if err != nil {
		slog.Error("app - Run - httpServer.Notify: " + err.Error())
	}
}
