package main

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/vladnishe/url-shortener/internal/config"
	database "github.com/vladnishe/url-shortener/internal/db"
	"github.com/vladnishe/url-shortener/internal/logger"
	"go.uber.org/zap"
)

const defaultCtxTimeout = 10 * time.Second

func main() {
	cfg, err := config.Load()
	if err != nil {
		panic(err)
	}

	log := logger.NewLogger(cfg.Env)
	defer log.Sync()

	db, err := database.NewPOSTGRES(cfg.DBUrl)
	if err != nil {
		log.Errorf("failed to initialize db: %v", err)
	}

	_ = db

	//r := router.NewRouter()

	srv := &http.Server{
		Addr:         fmt.Sprintf("%s:%d", cfg.Host, cfg.Port),
		Handler:      nil,
		WriteTimeout: cfg.Timeout,
		ReadTimeout:  cfg.Timeout,
		IdleTimeout:  cfg.IdleTimeout,
	}

	go func() {
		log.Info("starting server...", zap.Int64("port", cfg.Port))
		err := srv.ListenAndServe()
		if err != nil && !errors.Is(err, http.ErrServerClosed) {
			panic(err)
		}
	}()

	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs, syscall.SIGTERM, syscall.SIGINT)
	<-sigs
	log.Debug("starting shutting down server...")

	ctx, cancel := context.WithTimeout(context.Background(), defaultCtxTimeout)
	defer cancel()

	if err := srv.Shutdown(ctx); err != nil {
		fmt.Printf("failed to gracefully shutdown, making force shutdown")
	}

	log.Info("server stopped!")
}
