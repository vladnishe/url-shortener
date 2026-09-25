package logger

import (
	"go.uber.org/zap"
)

const (
	dev  = "dev"
	prod = "prod"
)

func NewLogger(env string) *zap.SugaredLogger {
	var logger *zap.SugaredLogger

	switch env {
	case dev:
		logger = zap.Must(zap.NewDevelopment()).Sugar()
	case prod:
		logger = zap.Must(zap.NewProduction()).Sugar()
	}

	return logger
}
