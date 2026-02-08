package cmd

import (
	"ddr13/riddlercore/internal/config"
	"ddr13/riddlercore/internal/conn"
	"ddr13/riddlercore/internal/queue"
	"log/slog"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"

	"github.com/spf13/cobra"
)

var (
	workerCmd = &cobra.Command{
		Use:   "worker",
		Short: "worker run rabbitmq worker",
		Long:  `worker run rabbitmq worker`,
		PreRun: func(cmd *cobra.Command, args []string) {

			slog.Info("Connecting database")
			if err := conn.ConnectDB(); err != nil {
				slog.Error("database connection failed", "error", err)
			}
			slog.Info("Database connected successfully!")

			slog.Info("Connecting rabbitmq")
			if err := conn.ConnectMachinery(config.MQ()); err != nil {
				slog.Error("rabbitmq connection failed", "error", err)
			}

			slog.Info("Connecting redis")
			if err := conn.ConnectDefaultRedis(); err != nil {
				slog.Error("redis connection failed", "error", err)
			}
			slog.Info("Redis connected successfully!")

			slog.Info("Initializing http-client")
			conn.InitClient()
			slog.Info("Initialized http-client successfully!")
		},
		Run: func(cmd *cobra.Command, args []string) {
			r := chi.NewRouter()
			// middlewares
			r.Use(middleware.Heartbeat("/"))

			db := conn.DefaultDB()

			mq := queue.New(conn.GetMachinery())
			_ = mq
			_ = db

			if err := conn.GetMachineryWorker().Launch(); err != nil {
				slog.Error("worker launch failed", "error", err)
			}
		},
	}
)

func init() {
	rootCmd.AddCommand(workerCmd)
}
