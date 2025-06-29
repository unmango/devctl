package cmd

import "github.com/spf13/cobra"

type InstallOptions struct {
	Method string
}

func NewInstall(options *InstallOptions) *cobra.Command {
	return &cobra.Command{
		Use:   "install",
		Short: "Install tools in the local workspace",
		Run: func(cmd *cobra.Command, args []string) {
		},
	}
}
