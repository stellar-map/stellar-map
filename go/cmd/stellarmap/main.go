package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

func main() {
	cfg, err := loadConfig()
	if err != nil {
		fmt.Println("unable to read config:", err)
		os.Exit(1)
	}

	rootCmd := &cobra.Command{
		Use:   "stellarmap",
		Short: "stellarmap is a blockchain and order book explorer for Stellar",
		Long:  "stellarmap is a blockchain and order book explorer for Stellar",
	}

	rootCmd.AddCommand(
		serverCommand(),
		ingestCommand(cfg),
	)

	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func loadConfig() (*viper.Viper, error) {
	cfg := viper.New()
	cfg.AutomaticEnv()
	cfg.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))
	cfg.AddConfigPath("$GOPATH/src/github.com/stellar-map/stellar-map/go/config")
	cfg.SetConfigName("config")
	cfg.SetConfigType("yaml")
	if err := cfg.ReadInConfig(); err != nil {
		return nil, err
	}
	return cfg, nil
}

func showErr(err error) {
	if err != nil {
		_, _ = fmt.Fprintf(os.Stderr, "%v\n", err)
		os.Exit(1)
	}
}
