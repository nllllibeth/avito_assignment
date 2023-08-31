package main

import (
	"log/slog"
	"net/http"
	"os"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	_ "github.com/lib/pq"

	"segmentation_service/internal/config"
	"segmentation_service/internal/http-server/handlers/segments"
	"segmentation_service/internal/http-server/handlers/users"
	"segmentation_service/internal/storage/postgres"
)

const (
	envLocal = "local"
	envDev   = "dev"
	envProd  = "prod"
) 

func main() {
	cfg := config.MustLoad()

	log := setupLogger(cfg.Env)

	log.Info("Welcome to User Segmentation API Service", slog.String("env", cfg.Env))
	log.Debug("Debug messages are enabled")

	storage, err := postgres.NewDB(postgres.Config{
		Host:     cfg.Host,
		Port:     cfg.Port,
		Username: cfg.Username,
		DBName:   cfg.DBName,
		SSLMode:  cfg.SSLMode,
		Password: cfg.Password,
	})

	if err != nil {
		log.Error("failed to init storage", err.Error())
		os.Exit(1)
	}
	router := chi.NewRouter()

	router.Use(middleware.RequestID)
	router.Use(middleware.Logger)
	router.Use(middleware.Recoverer)

	router.Post("/segment/create", segments.Create(log, storage))
	router.Delete("/segment/delete", segments.Delete(log, storage))
	router.Post("/user/add_segments", users.AddSegment(log, storage))
	router.Get("/user/get_segments", users.GetSegments(log, storage))

	log.Info("starting server", slog.String("address", cfg.Address))

	srv := &http.Server{
		Addr:         cfg.Address,
		Handler:      router,
		ReadTimeout:  cfg.HTTPServer.Timeout,
		WriteTimeout: cfg.HTTPServer.Timeout,
		IdleTimeout:  cfg.HTTPServer.IdleTimeout,
	}

	if err := srv.ListenAndServe(); err != nil {
		log.Error("failed to start server")
	}
	log.Error("server stopped")
}


func setupLogger(env string) *slog.Logger {
	var log *slog.Logger

	switch env {
	case envLocal:
		log = slog.New(
			slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelDebug}),
		)
	case envDev:
		log = slog.New(
			slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelDebug}),
		)
	case envProd:
		log = slog.New(
			slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelInfo}),
		)
	}
	return log
}
