package cmd

import (
	"github.com/spf13/cobra"
	"github.com/unmango/devctl/pkg/tool"
	"github.com/unmango/devctl/pkg/work"
	"github.com/unmango/go/cli"
)

var InstallCmd = NewInstall()

func NewInstall() *cobra.Command {
	return &cobra.Command{
		Use:   "install <name> <url>",
		Short: "Install tools in the local workspace",
		Args:  cobra.ExactArgs(2),
		Run: func(cmd *cobra.Command, args []string) {
			ctx := cmd.Context()
			tool, err := tool.FromConfig(args[0], tool.Config{
				Url: args[1],
			})
			if err != nil {
				cli.Fail(err)
			}

			work, err := work.Load(ctx)
			if err != nil {
				cli.Fail(err)
			}

			bin, err := work.Join("bin")
			if err != nil {
				cli.Fail(err)
			}

			if err = tool.Install(ctx, bin); err != nil {
				cli.Fail(err)
			}
		},
	}
}
