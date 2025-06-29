package config

import "github.com/spf13/cobra"

var Cmd = New()

func init() {
	Cmd.AddCommand(
		VerifyCmd,
	)
}

func New() *cobra.Command {
	return &cobra.Command{
		Use:   "config",
		Short: "Commands for managing configuration",
	}
}
