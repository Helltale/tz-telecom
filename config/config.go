package config

import (
	"fmt"

	"github.com/caarlos0/env/v9"
	"github.com/go-playground/validator/v10"
)

type Config struct {
	AppPort             string `env:"APP_PORT" envDefault:"8080" validate:"required,numeric"`
	DBHost              string `env:"DB_HOST" validate:"required,hostname|ip"`
	DBPort              string `env:"DB_PORT" envDefault:"5432" validate:"required,numeric"`
	DBUser              string `env:"DB_USER" validate:"required"`
	DBPass              string `env:"DB_PASS" validate:"required"`
	DBName              string `env:"DB_NAME" validate:"required"`
	DBSSLMode           string `env:"DB_SSLMODE" envDefault:"disable" validate:"required,oneof=disable require"`
	ReadTimeout         int    `env:"APP_READ_TIMEOUT" envDefault:"10" validate:"required,gte=1"`
	WriteTimeout        int    `env:"APP_WRITE_TIMEOUT" envDefault:"10" validate:"required,gte=1"`
	IdleTimeout         int    `env:"APP_IDLE_TIMEOUT" envDefault:"120" validate:"required,gte=10"`
	WorkerQueueLen      int    `env:"APP_WORKER_QUEUE_LEN" envDefault:"100" validate:"required,gte=1"`
	DBRetryInitialDelay int    `env:"DB_RETRY_INITIAL_DELAY" envDefault:"1" validate:"required,gte=1"` // сек
	DBRetryMaxDelay     int    `env:"DB_RETRY_MAX_DELAY" envDefault:"10" validate:"required,gte=1"`    // сек
	DBRetryMultiplier   int    `env:"DB_RETRY_MULTIPLIER" envDefault:"2" validate:"required,gte=1"`
	DBRetryMaxAttempts  int    `env:"DB_RETRY_MAX_ATTEMPTS" envDefault:"5" validate:"required,gte=1"`
}

var Conf Config
var validate *validator.Validate

func init() {
	validate = validator.New()
}

func Load() (*Config, error) {
	if err := env.Parse(&Conf); err != nil {
		return nil, fmt.Errorf("failed to load env vars: %w", err)
	}
	if err := validate.Struct(Conf); err != nil {
		return nil, fmt.Errorf("config validation error: %w", err)
	}
	return &Conf, nil
}
