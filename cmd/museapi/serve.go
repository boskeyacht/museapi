package museapi

import (
	"context"
	"database/sql"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/boskeyacht/museapi/server"
	"github.com/spf13/cobra"

	"github.com/uptrace/bun"
	"github.com/uptrace/bun/dialect/pgdialect"
	"github.com/uptrace/bun/driver/pgdriver"
	"github.com/uptrace/bun/extra/bundebug"
)

var (
	serveCmd = &cobra.Command{
		Use:   "serve",
		Short: "museapi server",
		Long:  `muse api server`,
		Run: func(cmd *cobra.Command, args []string) {
			sqldb := sql.OpenDB(pgdriver.NewConnector(pgdriver.WithDSN(cfg.PostgresURI)))
			db := bun.NewDB(sqldb, pgdialect.New())

			db.AddQueryHook(bundebug.NewQueryHook(
				bundebug.WithVerbose(true),
				bundebug.FromEnv("BUNDEBUG"),
			))

			server := server.NewServer(db, cfg)

			err := server.InitTables(context.Background())
			if err != nil {
				log.Fatalf("failed to initialize tables: %v", err)

				return
			}

			go server.InitRoutes(context.Background()).Run(":8080")

			log.Printf("\x1b[33m%s\x1b[0m", "Server started, interrupt to abort.")

			signal_channel := make(chan os.Signal, 1)
			signal.Notify(signal_channel, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)
			log.Printf("Exiting on signal '%v'", <-signal_channel)
		},
	}
)

func init() {
	rootCmd.AddCommand(serveCmd)
}
