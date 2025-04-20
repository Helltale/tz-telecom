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

	"github.com/Helltale/tz-telecom/config"
	"github.com/Helltale/tz-telecom/internal/delivery/httpdelivery"
	"github.com/Helltale/tz-telecom/internal/repository/postgresrepo"
	"github.com/Helltale/tz-telecom/internal/usecase"
	"github.com/Helltale/tz-telecom/internal/utils"
	"github.com/spf13/cobra"
)

var serveCmd = &cobra.Command{
	Use:   "serve",
	Short: "Run HTTP server",
	Run: func(cmd *cobra.Command, args []string) {
		cfg, err := config.Load()
		if err != nil {
			log.Fatalf("config error: %v", err)
		}

		db, err := utils.ConnectWithRetry(cfg, func(cfg *config.Config) (*sql.DB, error) {
			return postgresrepo.NewPostgres(cfg.BuildPostgresDSN())
		})
		if err != nil {
			log.Fatalf("DB connection failed: %v", err)
		}
		defer db.Close()

		userRepo := postgresrepo.NewUserRepo(db)
		orderRepo := postgresrepo.NewOrderRepo(db)

		userUC := usecase.NewUserUseCase(userRepo)
		orderUC := usecase.NewOrderUseCase(orderRepo)
		orderWorker := usecase.NewOrderWorker(orderUC, cfg.WorkerQueueLen)

		router := httpdelivery.NewRouter(userUC, orderWorker)

		srv := &http.Server{
			Addr:         ":" + cfg.AppPort,
			Handler:      router,
			ReadTimeout:  cfg.GetReadTimeout(),
			WriteTimeout: cfg.GetWriteTimeout(),
			IdleTimeout:  cfg.GetIdleTimeout(),
		}

		go func() {
			log.Printf("Listening on %s", cfg.AppPort)
			if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
				log.Fatalf("ListenAndServe: %v", err)
			}
		}()

		quit := make(chan os.Signal, 1)
		signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
		<-quit

		ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()
		srv.Shutdown(ctx)
		log.Println("Server gracefully stopped")
	},
}

func init() {
	rootCmd.AddCommand(serveCmd)
}
