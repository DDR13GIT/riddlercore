package cmd

import (
	"ddr13/riddlercore/internal/config"
	"log/slog"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var (
	cfgFile                 string
	verbose, prettyPrintLog bool

	rootCmd = &cobra.Command{
		Use:   "riddlercore",
		Short: "riddlercore provides riddlercore related core business logic",
		Long:  `riddlercore provides riddlercore related core business logic`,
	}
)

func init() {
	cobra.OnInitialize(initConfig)
	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "config.yaml", "config file")
	rootCmd.PersistentFlags().BoolVarP(&prettyPrintLog, "pretty", "p", false, "pretty print verbose/slog")
	rootCmd.PersistentFlags().BoolVarP(&verbose, "verbose", "v", false, "verbose output")

	viper.BindPFlag("verbose", rootCmd.PersistentFlags().Lookup("verbose"))

}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		slog.Error("root command execution failed", "error", err)
	}
}

func initConfig() {
	slog.Info("Loading configurations")
	config.Init()
	slog.Info("Configurations loaded successfully!")
}
