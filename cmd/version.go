package cmd

import (
	"os"

	"github.com/spf13/cobra"
	"github.com/unmango/devctl/pkg/version"
	"github.com/unmango/devctl/pkg/version/opts"
	"github.com/unmango/devctl/pkg/work"
	"github.com/unmango/go/cli"
)

var VersionCmd = NewVersion()

type VersionOptions struct {
	work.ChdirOptions
	Prefixed bool
}

func (o VersionOptions) PrintPrefixed() opts.PrintOp {
	return opts.PrintPrefixed(o.Prefixed)
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

			_, err := version.Cat(
				version.RelPath(args[0]),
				os.Stdout,
				opts.PrintPrefixed(),
			)
			if err != nil {
				cli.Fail(err)
			}
		},
	}

	_ = work.ChdirFlag(cmd, &opts.ChdirOptions, "")
	cmd.Flags().BoolVarP(&opts.Prefixed, "prefixed", "p", false,
		`include a leading 'v' in the version`,
	)

	return cmd
}
