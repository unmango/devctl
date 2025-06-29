package cmd

import "github.com/spf13/cobra"

func NewInstall() *cobra.Command {
	return &cobra.Command{
		Use: "install",
	}
}
