package main

import (
	"fmt"
	"os"

	"github.com/charmbracelet/log"
	"github.com/spf13/cobra"
	"github.com/unmango/devctl/pkg/cmd"
)

var rootCmd = &cobra.Command{
	Use:   "devctl",
	Short: "Helper utilities for developing code",
}

func main() {
	log.SetLevel(log.ErrorLevel)

	rootCmd.AddCommand(
		cmd.NewInit(),
		cmd.NewList(&cmd.ListOptions{}),
		cmd.NewVersion(),
	)

	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
