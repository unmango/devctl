package init

import (
	"github.com/spf13/cobra"
)

var Cmd = New()

func init() {
	Cmd.AddCommand(
		ConfigCmd,
		VersionCmd,
	)
}

type InitOptions struct{}

func New() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "init [scaffold]",
		Short: "Generates files for the specified scaffold",
	}

	return cmd
}
