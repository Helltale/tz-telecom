package cmd

import (
	"log"

	"github.com/Helltale/tz-telecom/config"
	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"github.com/spf13/cobra"
)

var migrateCmd = &cobra.Command{
	Use:   "migrate",
	Short: "Run DB migrations",
	Run: func(cmd *cobra.Command, args []string) {
		cfg, err := config.Load()
		if err != nil {
			log.Fatalf("failed to load config: %v", err)
		}

		m, err := migrate.New(
			"file://internal/repository/postgresrepo/migrations",
			cfg.BuildPostgresDSN(),
		)
		if err != nil {
			log.Fatalf("failed to create migration instance: %v", err)
		}

		if err := m.Up(); err != nil && err.Error() != "no change" {
			log.Fatalf("migration error: %v", err)
		}

		log.Println("migrations applied successfully")
	},
}

func init() {
	rootCmd.AddCommand(migrateCmd)
}
