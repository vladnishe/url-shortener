package main

import (
	"fmt"
	"net/http"

	"github.com/vladnishe/url-shortener/internal/config"
	"github.com/vladnishe/url-shortener/internal/router"
)

func main() {
	cfg, err := config.MustLoad()
	if err != nil {
		panic(err)
	}
	fmt.Println(cfg)

	r := router.NewRouter()

	srv := &http.Server{
		Addr: fmt.Sprintf("%s:%d", cfg.Host, cfg.Port),
		Handler: r,
		WriteTimeout: cfg.Timeout,
		ReadTimeout: cfg.Timeout,
		IdleTimeout: cfg.IdleTimeout,
	}

	if err := srv.ListenAndServe(); err != nil {
		panic(err)
	}
}
