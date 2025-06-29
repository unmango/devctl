package config

import (
	"os"

	"github.com/charmbracelet/log"
	"github.com/spf13/cobra"
	"github.com/unmango/devctl/pkg/config"
	"github.com/unmango/devctl/pkg/work"
	"github.com/unmango/go/cli"
)

var VerifyCmd = NewVerify()

func NewVerify() *cobra.Command {
	return &cobra.Command{
		Use:   "verify",
		Short: "Verify the current configuration",
		Run: func(cmd *cobra.Command, args []string) {
			log := log.NewWithOptions(os.Stdout, log.Options{
				Level:           log.InfoLevel,
				ReportTimestamp: false,
			})

			log.Info("Loading workspace")
			work, err := work.Load(cmd.Context())
			if err != nil {
				cli.Fail(err)
			}

			log.Info("Loading config")
			c, err := config.FromDirectory(work)
			if err != nil {
				cli.Fail(err)
			}

			if c != nil {
				log.Info("Configuration is valid!")
			} else {
				log.Info("No configuration to verify")
			}
		},
	}
}
