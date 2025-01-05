package init

import (
	"fmt"
	"strings"

	"github.com/spf13/cobra"
	"github.com/unmango/devctl/pkg/version"
	"github.com/unmango/devctl/pkg/work"
	"github.com/unmango/go/cli"
)

var (
	AutoVersionSource   = "auto"
	GitHubVersionSource = "github"

	VersionSources = []string{
		AutoVersionSource,
		GitHubVersionSource,
	}
)

type VersionOptions struct {
	work.ChdirOptions
	Name     string
	Source   string
	Makefile bool
}

func NewVersion() *cobra.Command {
	opts := VersionOptions{}

	cmd := &cobra.Command{
		Use: "version [dependency]",
		Example: `devctl init version my-app v0.0.69
devctl init version 0.0.69 --name my-app
devctl init version my-app 0.0.69 --makefile`,
		Short:   "Generates files for versioning the specified dependency",
		Aliases: []string{"v"},
		Args:    cobra.RangeArgs(1, 2),
		Run: func(cmd *cobra.Command, args []string) {
			ctx := cmd.Context()
			if err := opts.Chdir(ctx); err != nil {
				cli.Fail(err)
			}

			var (
				dep  = args[len(args)-1]
				name string
				src  version.Source
				err  error
			)

			if len(args) == 2 {
				name = args[0]
			} else {
				name = opts.Name
			}

			switch opts.Source {
			case AutoVersionSource:
				src, err = version.GuessSource(dep)
			case GitHubVersionSource:
				src = version.GitHub(dep)
			default:
				err = fmt.Errorf("unsupported source: %s", opts.Source)
			}
			if err != nil {
				cli.Fail(err)
			}

			if err = version.Init(ctx, name, src); err != nil {
				cli.Fail(err)
			}
			if opts.Makefile {
				if err = version.WriteMakefile(name); err != nil {
					cli.Fail(err)
				}
			}
		},
	}

	_ = work.ChdirFlag(cmd, &opts.ChdirOptions, "")
	cmd.Flags().StringVarP(&opts.Source, "source", "s", AutoVersionSource,
		fmt.Sprintf("source of [dependency]: %s", strings.Join(VersionSources, ", ")),
	)
	cmd.Flags().StringVarP(&opts.Name, "name", "n", "", "explicit dependency name")
	cmd.Flags().BoolVar(&opts.Makefile, "makefile", false, "generate a Makefile variable")

	return cmd
}
