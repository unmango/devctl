package cmd

import (
	"github.com/charmbracelet/log"
	"github.com/spf13/cobra"
	"github.com/unmango/devctl/cmd/config"
	"github.com/unmango/devctl/cmd/initialize"
	"github.com/unmango/devctl/cmd/migrate"
	"github.com/unmango/devctl/cmd/stack"
)

var root = &cobra.Command{
	Use:   "devctl [path]",
	Short: "Helper utilities for developing code",
}

func init() {
	root.AddCommand(
		initialize.Cmd,
		config.Cmd,
		InstallCmd,
		ListCmd,
		LocalBinCmd,
		migrate.Cmd,
		stack.Cmd,
		VersionCmd,
	)
}

func Execute() error {
	log.SetReportTimestamp(false)
	return root.Execute()
}
