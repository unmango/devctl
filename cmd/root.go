package cmd

import (
	"github.com/charmbracelet/log"
	"github.com/spf13/cobra"
	"github.com/unmango/devctl/cmd/config"
	"github.com/unmango/devctl/cmd/initialize"
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
		VersionCmd,
	)
}

func Execute() error {
	log.SetReportTimestamp(false)
	return root.Execute()
}
