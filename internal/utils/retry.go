package utils

import (
	"fmt"
	"log/slog"
	"time"

	"github.com/Helltale/tz-telecom/config"
)

func ConnectWithRetry[T any](cfg *config.Config, connectFunc func(*config.Config) (T, error)) (T, error) {
	var result T
	var err error

	delay := time.Duration(cfg.DBRetryInitialDelay) * time.Second
	maxDelay := time.Duration(cfg.DBRetryMaxDelay) * time.Second
	multiplier := time.Duration(cfg.DBRetryMultiplier)
	attempts := cfg.DBRetryMaxAttempts

	for i := 0; i < attempts; i++ {
		result, err = connectFunc(cfg)
		if err == nil {
			return result, nil
		}

		slog.Warn("cannot connect to database",
			slog.Int("attempt", i+1),
			slog.String("error", err.Error()),
		)

		time.Sleep(delay)
		if delay < maxDelay {
			delay *= multiplier
			if delay > maxDelay {
				delay = maxDelay
			}
		}
	}

	return result, fmt.Errorf("failed to connect after %d attempts: %w", attempts, err)
}
