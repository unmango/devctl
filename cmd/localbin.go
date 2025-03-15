package cmd

import (
	"fmt"
	"path/filepath"

	"github.com/spf13/cobra"
	"github.com/unmango/devctl/pkg/work"
	"github.com/unmango/go/cli"
)

func NewLocalBin() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "localbin",
		Short: "Prints the absolute path to the workspace's binary directory",
		Run: func(cmd *cobra.Command, args []string) {
			if work, err := work.Load(cmd.Context()); err != nil {
				cli.Fail(err)
			} else {
				fmt.Println(filepath.Join(work.Path(), "bin"))
			}
		},
	}

	return cmd
}
