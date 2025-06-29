package cmd

import (
	"github.com/charmbracelet/log"
	"github.com/spf13/cobra"
	initcmd "github.com/unmango/devctl/cmd/init"
)

var root = &cobra.Command{
	Use:   "devctl [path]",
	Short: "Helper utilities for developing code",
}

func init() {
	root.AddCommand(
		initcmd.Cmd,
		ListCmd,
		LocalBinCmd,
		VersionCmd,
	)
}

func Execute() error {
	log.SetReportTimestamp(false)
	return root.Execute()
}
