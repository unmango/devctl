package cmd

import (
	"github.com/charmbracelet/log"
	"github.com/spf13/cobra"
	"github.com/unmango/devctl/pkg/config"
	"github.com/unmango/devctl/pkg/tool"
	"github.com/unmango/devctl/pkg/work"
	"github.com/unmango/go/cli"
)

var InstallCmd = NewInstall()

func NewInstall() *cobra.Command {
	return &cobra.Command{
		Use:   "install",
		Short: "Install tools in the local workspace",
		Run: func(cmd *cobra.Command, args []string) {
			log.SetLevel(log.InfoLevel)

			work, err := work.Load(cmd.Context())
			if err != nil {
				cli.Fail(err)
			}

			config, err := config.FromDirectory(work)
			if err != nil {
				cli.Fail(err)
			}

			for n, c := range config.Tools {
				if t, err := tool.FromConfig(n, c); err != nil {
					cli.Fail(err)
				} else {
					log.Infof("Loaded tool: %s", t.Name)
				}
			}
		},
	}
}
