package initialize

import (
	"github.com/spf13/cobra"
	"github.com/unmango/devctl/pkg/config"
	"github.com/unmango/devctl/pkg/work"
	"github.com/unmango/go/cli"
)

var ConfigCmd = NewConfig()

func NewConfig() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "config",
		Short: "Initializes a new config file",
		Run: func(cmd *cobra.Command, args []string) {
			work, err := work.Load(cmd.Context())
			if err != nil {
				cli.Fail(err)
			}

			if err = config.Init(work); err != nil {
				cli.Fail(err)
			}
		},
	}

	return cmd
}
