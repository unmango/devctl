package main

import (
	"os"

	"github.com/charmbracelet/log"
	"github.com/spf13/cobra"
	"github.com/unmango/devctl/cmd"
	"github.com/unmango/devctl/pkg/list"
	"github.com/unmango/devctl/pkg/version"
	"github.com/unmango/go/cli"
)

var rootCmd = &cobra.Command{
	Use:   "devctl [path]",
	Short: "Helper utilities for developing code",
}

func main() {
	log.SetLevel(log.ErrorLevel)

	rootCmd.AddCommand(
		cmd.NewInit(),
		cmd.NewList(&list.Options{}),
		cmd.NewLocalBin(),
		cmd.NewVersion(),
	)

	if len(os.Args) == 2 {
		n, err := version.PrintIfPath(os.Args[1])
		if err != nil {
			cli.Fail(err)
		}
		if n > 0 {
			return
		}
	}

	if err := rootCmd.Execute(); err != nil {
		cli.Fail(err)
	}
}
