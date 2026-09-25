package main

import (
	"fmt"

	"github.com/vladnishe/url-shortener/internal/config"
)

func main() {
	cfg, err := config.MustLoad()
	if err != nil {
		panic(err)
	}
	fmt.Println(cfg)
}
