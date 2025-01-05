package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/unmango/devctl/pkg/version"
	"github.com/unmango/devctl/pkg/work"
	"github.com/unmango/go/cli"
)

type VersionOptions struct {
	work.ChdirOptions
}

func NewVersion() *cobra.Command {
	opts := VersionOptions{}

	cmd := &cobra.Command{
		Use:     "version [name]",
		Short:   "Print the version of the specified dependency",
		Aliases: []string{"v"},
		Args:    cobra.ExactArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			if err := opts.Chdir(cmd.Context()); err != nil {
				cli.Fail(err)
			}

			v, err := version.Read(args[0])
			if err != nil {
				cli.Fail(err)
			}

			fmt.Println(v)
		},
	}

	_ = work.ChdirFlag(cmd, &opts.ChdirOptions, "")

	return cmd
}
