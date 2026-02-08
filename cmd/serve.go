package cmd

import (
	"context"
	"fmt"
	"log/slog"
	"net/http"
	"os"
	"os/signal"
	"time"

	"ddr13/riddlercore/internal/config"
	"ddr13/riddlercore/internal/conn"
	"ddr13/riddlercore/internal/utils"
	_advertisementHttp "ddr13/riddlercore/questionengine/delivery/http"
	_advertisementRepository "ddr13/riddlercore/questionengine/repository"
	_advertisementUsecase "ddr13/riddlercore/questionengine/usecase"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	validator "github.com/go-playground/validator/v10"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"github.com/spf13/cobra"
)

var (
	serveCmd = &cobra.Command{
		Use:   "serve",
		Short: "Serve run available servers such as: HTTP/JSON or gRPC",
		Long:  `Serve run available servers such as: HTTP/JSON or gRPC`,
		PreRun: func(cmd *cobra.Command, args []string) {

			slog.Info("Connecting database")
			if err := conn.ConnectDB(); err != nil {
				slog.Error("database connection failed", "error", err)
			}
			slog.Info("Database connected successfully!")

			slog.Info("Connecting rabbitmq")
			if err := conn.ConnectMachinery(config.MQ()); err != nil {
				slog.Error("rabbitMQ connection failed", "error", err)
			}
			slog.Info("Rabbitmq connected successfully!")

			slog.Info("Connecting redis")
			if err := conn.ConnectDefaultRedis(); err != nil {
				slog.Error("redis connection failed", "error", err)
			}
			slog.Info("Redis connected successfully!")

			slog.Info("Initializing http-client")
			conn.InitClient()
			slog.Info("Initialized http-client successfully!")
		},
		Run: serve,
	}
)

func init() {
	rootCmd.AddCommand(serveCmd)
}

func serve(cmd *cobra.Command, args []string) {
	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt)

	srv := buildHTTP(cmd, args)

	go func(sr *http.Server) {
		if err := sr.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			slog.Error("listen and serve failed", "error", err)
		}
	}(srv)

	<-stop
	slog.Info("Shutting down HTTP server")
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := srv.Shutdown(ctx); err != nil && err != http.ErrServerClosed {
		slog.Error("http server closed", "error", err)
	}

	slog.Info("server shutdown successful!")
}

func buildHTTP(_ *cobra.Command, _ []string) *http.Server {
	r := chi.NewRouter()
	r.Use(middleware.RequestID)
	r.Use(middleware.Logger)
	r.Use(middleware.RealIP)
	r.Use(middleware.Recoverer)
	r.Use(middleware.Heartbeat("/"))

	r.Mount("/metrics", promhttp.Handler())

	cfg := config.App()
	db := conn.DefaultDB()
	_ = db

	advertisementRepo := _advertisementRepository.New(db)
	advertisementUsecase := _advertisementUsecase.New(advertisementRepo)
	validate := validator.New(validator.WithRequiredStructEnabled())
	_ = validate.RegisterValidation("future", utils.FutureValidator)
	_ = validate.RegisterValidation("notblank", utils.NotBlankValidator)
	_advertisementHttp.NewHTTPHandler(r, advertisementUsecase, validate)

	slog.Info("HTTP Listening on port: ", cfg.HTTPPort)
	slog.Info("For system check use cURL request: ", "[curl localhost:"+fmt.Sprint(cfg.HTTPPort))

	return &http.Server{
		Addr:              fmt.Sprintf(":%d", cfg.HTTPPort),
		Handler:           r,
		ReadHeaderTimeout: cfg.ReadTimeout,
		WriteTimeout:      cfg.WriteTimeout,
		IdleTimeout:       cfg.IdleTimeout,
	}
}
