package config

import (
	"fmt"
	"os"
	"strconv"
	"time"

	"github.com/joho/godotenv"
)

const (
	defaultEnv               = "dev"
	defaultHost              = "localhost"
	defaultPort        int64 = 8000
	defaultIdleTimeout       = 60 * time.Second
	defaultTimeout           = 30 * time.Second
)

type Config struct {
	Env         string
	Host        string
	Port        int64
	IdleTimeout time.Duration
	Timeout     time.Duration
	DBConfig
}

func Load() (*Config, error) {
	_ = godotenv.Load()

	dbCfg, err := loadDBConfig()
	if err != nil {
		return nil, fmt.Errorf("failed to set database configuration: %v", err)
	}

	return &Config{
		Env:         getEnv("ENV", defaultEnv),
		Host:        getEnv("HOST", defaultHost),
		Port:        getEnvAsInt64("PORT", defaultPort),
		DBConfig:    dbCfg,
		IdleTimeout: getEnvAsDuration("IDLE_TIMEOUT", defaultIdleTimeout),
		Timeout:     getEnvAsDuration("TIMEOUT", defaultTimeout),
	}, nil
}

func getEnv(key, fallback string) string {
	val, exists := os.LookupEnv(key)
	if !exists {
		return fallback
	}
	return val
}

func getEnvAsInt64(key string, fallback int64) int64 {
	val, exists := os.LookupEnv(key)
	if !exists {
		return fallback
	}

	//Return fallback, no need to work with error at this point
	valInt64, err := strconv.ParseInt(val, 10, 64)
	if err != nil {
		return fallback
	}

	return valInt64
}

func getEnvAsDuration(key string, fallback time.Duration) time.Duration {
	val, exists := os.LookupEnv(key)
	if !exists {
		return fallback
	}

	valDuration, err := time.ParseDuration(val)
	if err != nil {
		return fallback
	}
	return valDuration
}
