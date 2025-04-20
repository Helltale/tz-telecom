package cmd

import (
	"context"
	"database/sql"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/getsentry/sentry-go"
	"github.com/spf13/cobra"

	"github.com/Helltale/tz-telecom/config"
	"github.com/Helltale/tz-telecom/internal/delivery/httpdelivery"
	"github.com/Helltale/tz-telecom/internal/repository/postgresrepo"
	"github.com/Helltale/tz-telecom/internal/usecase"
	"github.com/Helltale/tz-telecom/internal/utils"
)

var serveCmd = &cobra.Command{
	Use:   "serve",
	Short: "Run HTTP server",
	Run: func(cmd *cobra.Command, args []string) {

		cfg, err := config.Load()
		if err != nil {
			log.Fatalf("config error: %v", err)
		}

		// initialize sentry if DSN is provided
		if cfg.SentryDSN != "" {
			err := sentry.Init(sentry.ClientOptions{
				Dsn:              cfg.SentryDSN,
				Environment:      cfg.SentryEnv,
				TracesSampleRate: cfg.SentrySampleRate,
			})
			if err != nil {
				log.Fatalf("sentry init failed: %v", err)
			}
			defer sentry.Flush(2 * time.Second)
			log.Println("Sentry initialized")
		}

		db, err := utils.ConnectWithRetry(cfg, func(cfg *config.Config) (*sql.DB, error) {
			return postgresrepo.NewPostgres(cfg.BuildPostgresDSN())
		})
		if err != nil {
			log.Fatalf("DB connection failed: %v", err)
		}
		defer db.Close()

		// repositories
		userRepo := postgresrepo.NewUserRepo(db)
		orderRepo := postgresrepo.NewOrderRepo(db)

		// use cases
		userUC := usecase.NewUserUseCase(userRepo)
		orderUC := usecase.NewOrderUseCase(orderRepo)

		// worker
		orderWorker := usecase.NewOrderWorker(orderUC, cfg.WorkerQueueLen)

		// router
		router := httpdelivery.NewRouter(userUC, orderWorker)

		// server
		srv := &http.Server{
			Addr:         ":" + cfg.AppPort,
			Handler:      router,
			ReadTimeout:  cfg.GetReadTimeout(),
			WriteTimeout: cfg.GetWriteTimeout(),
			IdleTimeout:  cfg.GetIdleTimeout(),
		}

		// start server
		go func() {
			log.Printf("Listening on port %s...", cfg.AppPort)
			if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
				log.Fatalf("HTTP server error: %v", err)
			}
		}()

		quit := make(chan os.Signal, 1)
		signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
		<-quit

		ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()
		if err := srv.Shutdown(ctx); err != nil {
			log.Fatalf("Shutdown failed: %v", err)
		}
		log.Println("Server gracefully stopped")
	},
}

func init() {
	rootCmd.AddCommand(serveCmd)
}
