package cmd

import (
	"github.com/spf13/cobra"
	gen "github.com/unmango/devctl/pkg/cmd/init"
)

type InitOptions struct{}

func NewInit() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "init [scaffold]",
		Short: "Generates files the specified scaffold",
	}

	cmd.AddCommand(
		gen.NewVersion(),
	)

	return cmd
}
