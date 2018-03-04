package main

import (
	"fmt"

	"github.com/spf13/cobra"
)

func serverCommand() *cobra.Command {
	var (
		port int
	)

	cmd := &cobra.Command{
		Use:   "server",
		Short: "Start the HTTP server",
		Long:  "Start the HTTP server",
		PreRunE: func(cmd *cobra.Command, args []string) error {
			return nil
		},
		Run: func(cmd *cobra.Command, args []string) {
			serve()
		},
	}
	cmd.Flags().IntVar(&port, "port", 8080, `Port to listen on`)

	return cmd
}

func serve() {
	fmt.Println("running HTTP server")
}
