package config

import (
	"fmt"
	"os"
)

type DBConfig struct {
	name     string
	user     string
	password string
	host     string
	port     string
	sslMode  string
}

func loadDBConfig() (DBConfig, error) {
	host := os.Getenv("DB_HOST")
	if host == "" {
		return DBConfig{}, fmt.Errorf("DB_HOST is required")
	}
	name := os.Getenv("DB_NAME")
	if name == "" {
		return DBConfig{}, fmt.Errorf("DB_NAME is required")
	}
	user := os.Getenv("DB_USER")
	if user == "" {
		return DBConfig{}, fmt.Errorf("DB_USER is required")
	}
	password := os.Getenv("DB_PASSWORD")
	if password == "" {
		return DBConfig{}, fmt.Errorf("DB_PASSWORD is required")
	}

	return DBConfig{
		name:     name,
		user:     user,
		password: password,
		host:     host,
		port:     getEnv("DB_PORT", "5432"),
		sslMode:  getEnv("DB_SSL", "disable"),
	}, nil
}

func (c *DBConfig) DSN() string {
	return fmt.Sprintf(
		"postgres://%s:%s@%s:%s/%s?sslmode=%s",
		c.user,
		c.password,
		c.host,
		c.port,
		c.name,
		c.sslMode,
	)
}
