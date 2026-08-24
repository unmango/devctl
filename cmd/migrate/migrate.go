package migrate

import (
	"github.com/spf13/cobra"
)

var Cmd = New()

func init() {
	Cmd.AddCommand(
		AgentsCmd,
	)
}

func New() *cobra.Command {
	return &cobra.Command{
		Use:   "migrate",
		Short: "Migrate a repository between conventions",
	}
}
