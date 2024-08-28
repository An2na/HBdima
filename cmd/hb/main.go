package main

import (
	"context"
	"fmt"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/cors"
	"hbDima/internal/config"
	"hbDima/internal/http-server/handlers/greetingGenerate"
	"hbDima/internal/http-server/handlers/showPage"
	"hbDima/internal/lib/logger/sl"
	"hbDima/internal/storage/sqlite"
	"log/slog"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

const (
	envLocal = "local"
	envDev   = "dev"
	envProd  = "prod"
)

func main() {
	cfg := config.MustLoad()
	fmt.Println(cfg)

	log := setupLogger(cfg.Env)

	log.Info("starting hbDima", slog.String("env", cfg.Env)) // к каждому сообщению будет добавляться поле с информацией о текущем окружении
	log.Debug("debug messages are enabled")

	storage, err := sqlite.New(cfg.StoragePath)
	if err != nil {
		log.Error("failed to init storage", sl.Err(err))
		os.Exit(1) //return
	}

	congratulationsHandler := &greetingGenerate.CongratulationsHandler{Storage: storage}

	router := chi.NewRouter()
	router.Use(cors.Handler(cors.Options{
		AllowedOrigins:   []string{"https://*", "http://*", "*"}, // В продакшне лучше указать конкретные домены
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "Access-Control-Allow-Origin", "X-CSRF-Token"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: false,
		MaxAge:           300, // Максимальное значение, которое не игнорируется основными браузерами
	}))
	router.Use(middleware.RequestID) // Добавляет request_id в каждый запрос
	router.Use(middleware.Logger)    // Логирование всех запросов
	router.Use(middleware.Recoverer) // Восстановление после паники в приложении

	router.Get("/home", showPage.Home(log))
	router.Get("/bio", showPage.Bio(log))
	router.Get("/nastya", showPage.Nastya(log))
	router.Get("/mom", showPage.Mom(log))
	router.Get("/anya", showPage.Anya(log))
	router.Get("/cats", showPage.Cats(log))
	router.Get("/happy", showPage.Birth(log))
	router.Get("/foto", showPage.Foto(log))
	router.Get("/greetings", greetingGenerate.Congratulate(log, congratulationsHandler))
	router.Get("/congrats", greetingGenerate.GenerateGreetingHandler(log, congratulationsHandler))
	fs := http.FileServer(http.Dir("static-fiels/"))
	router.Handle("/static-fiels/*", http.StripPrefix("/static-fiels/", fs))

	done := make(chan os.Signal, 1)
	signal.Notify(done, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)

	srv := &http.Server{
		Addr:         cfg.Address,
		Handler:      router,
		ReadTimeout:  cfg.HTTPServer.Timeout,
		WriteTimeout: cfg.HTTPServer.Timeout,
		IdleTimeout:  cfg.HTTPServer.IdleTimeout,
	}
	go func() {
		if err := srv.ListenAndServe(); err != nil {
			log.Error("failed to start server")
		}
	}()

	log.Info("server started")

	<-done
	log.Info("stopping server")

	// TODO: move timeout to config
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	if err := srv.Shutdown(ctx); err != nil {
		log.Error("failed to stop server", sl.Err(err))
		return
	}

	// TODO: close storage

	log.Info("server stopped")
}

func setupLogger(env string) *slog.Logger {
	var log *slog.Logger
	fmt.Println("kek", env)

	switch env {
	case envLocal:
		log = slog.New(
			slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelDebug}),
		)
		//log = setupPrettySlog()
	case envDev:
		log = slog.New(
			slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelDebug}),
		)
	case envProd:
		log = slog.New(
			slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelInfo}),
		)
	default:
		log = slog.New(
			slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelInfo}),
		)
	}

	return log
}
