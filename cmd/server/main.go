package main

import (
	"context"
	"errors"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/vladnishe/url-shortener/internal/config"
	"github.com/vladnishe/url-shortener/internal/router"
)

const defaultCtxTimeout = 10 * time.Second

func main() {
	cfg, err := config.MustLoad()
	if err != nil {
		panic(err)
	}

	r := router.NewRouter()

	srv := &http.Server{
		Addr:         fmt.Sprintf("%s:%d", cfg.Host, cfg.Port),
		Handler:      r,
		WriteTimeout: cfg.Timeout,
		ReadTimeout:  cfg.Timeout,
		IdleTimeout:  cfg.IdleTimeout,
	}

	go func() {
		log.Printf("server started on port %d", cfg.Port)
		err := srv.ListenAndServe()
		if err != nil && !errors.Is(err, http.ErrServerClosed) {
			panic(err)
		}
	}()

	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs, syscall.SIGTERM, syscall.SIGINT)
	<-sigs
	log.Print("starting shutting down server...")

	ctx, cancel := context.WithTimeout(context.Background(), defaultCtxTimeout)
	defer cancel()

	if err := srv.Shutdown(ctx); err != nil {
		fmt.Printf("failed to gracefully shutdown, making force shutdown")
	}

	log.Print("server stopped!")
}
