package config

import "github.com/spf13/cobra"

var Cmd = New()

func New() *cobra.Command {
	return &cobra.Command{
		Use:   "config",
		Short: "Commands for managing configuration",
	}
}
