package main

import (
	"fmt"
	"os"

	"github.com/charmbracelet/log"
	"github.com/spf13/cobra"
	"github.com/unmango/devctl/pkg/cmd"
	"github.com/unmango/devctl/pkg/version"
	"github.com/unmango/go/cli"
)

var rootCmd = &cobra.Command{
	Use:   "devctl [path]",
	Short: "Helper utilities for developing code",
	Args:  cobra.RangeArgs(0, 1),
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 1 && version.IsPath(args[0]) {
			if _, err := version.Cat(args[0]); err != nil {
				cli.Fail(err)
			}
		} else {
			_ = cmd.Help()
		}
	},
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
