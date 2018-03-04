package main

import (
	"context"

	"github.com/pkg/errors"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"

	"github.com/stellar-map/stellar-map/go/pkg/db"
	"github.com/stellar-map/stellar-map/go/pkg/ingest"
)

func ingestCommand(cfg *viper.Viper) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "ingest",
		Short: "Ingest data from Horizon server",
		Long:  "Ingest data from Horizon server",
		PreRunE: func(cmd *cobra.Command, args []string) error {
			resource := cfg.GetString("ingest.resource")
			if resource != "ledger" && resource != "transaction" && resource != "payment" {
				return errors.New("resource must be one of ledger or transaction or payment")
			}
			return nil
		},
		Run: func(cmd *cobra.Command, args []string) {
			ingestMain(cfg)
		},
	}
	cmd.Flags().String("resource", "ledger", `resource to ingest (one of "ledger", "transaction", "payment")`)
	cmd.Flags().Bool("testnet", false, `ingest from the TestNet network`)
	cmd.Flags().String("cursor", "", `cursor to start from, "0" for beginning (required)`)
	cmd.MarkFlagRequired("cursor")

	cfg.BindPFlag("ingest.resource", cmd.Flags().Lookup("resource"))
	cfg.BindPFlag("ingest.testnet", cmd.Flags().Lookup("testnet"))
	cfg.BindPFlag("ingest.cursor", cmd.Flags().Lookup("cursor"))

	return cmd
}

func ingestMain(cfg *viper.Viper) {
	ctx := context.Background()

	repo, err := db.New(cfg.GetString("db.url"))
	if err != nil {
		showErr(err)
	}

	ingester := ingest.New(repo, cfg.GetBool("ingest.testnet"), true)
	resource := cfg.GetString("ingest.resource")
	cursor := cfg.GetString("ingest.cursor")

	switch resource {
	case "ledger":
		err = ingester.Ledgers(ctx, cursor)
	case "transaction":
		// err = client.StreamAllTransactions(ctx, &horizonCursor, handleTransaction)
	case "payment":
		// err = client.StreamPayments(ctx, &horizonCursor, handlePayment)
	default:
		err = errors.Errorf("unknown resource %s", resource)
	}

	showErr(err)
}
