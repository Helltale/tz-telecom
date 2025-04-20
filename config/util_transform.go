package config

import (
	"fmt"
	"time"
)

func (cfg Config) GetReadTimeout() time.Duration {
	return time.Duration(cfg.ReadTimeout) * time.Second
}

func (cfg Config) GetWriteTimeout() time.Duration {
	return time.Duration(cfg.WriteTimeout) * time.Second
}

func (cfg Config) GetIdleTimeout() time.Duration {
	return time.Duration(cfg.IdleTimeout) * time.Second
}

func (cfg Config) BuildPostgresDSN() string {
	return fmt.Sprintf(
		"postgres://%s:%s@%s:%s/%s?sslmode=%s",
		cfg.DBUser,
		cfg.DBPass,
		cfg.DBHost,
		cfg.DBPort,
		cfg.DBName,
		cfg.DBSSLMode,
	)
}
