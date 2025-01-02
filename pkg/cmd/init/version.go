package init

import (
	"fmt"
	"strings"

	"github.com/spf13/cobra"
	"github.com/unmango/devctl/pkg/version"
	"github.com/unmango/devctl/pkg/work"
	util "github.com/unmango/go/cmd"
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
	Name   string
	Source string
}

func NewVersion() *cobra.Command {
	opts := VersionOptions{}

	cmd := &cobra.Command{
		Use:   "version [dependency]",
		Short: "Generates files for versioning the specified dependency",
		Args:  cobra.RangeArgs(1, 2),
		Run: func(cmd *cobra.Command, args []string) {
			ctx := cmd.Context()
			if err := opts.Chdir(ctx); err != nil {
				util.Fail(err)
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
				util.Fail(err)
			}

			if err = version.Init(ctx, name, src); err != nil {
				util.Fail(err)
			}
		},
	}

	_ = work.ChdirFlag(cmd, &opts.ChdirOptions, "")
	cmd.Flags().StringVarP(&opts.Source, "source", "s", AutoVersionSource,
		fmt.Sprintf("source of dependency, one of: [%s]", strings.Join(VersionSources, ", ")),
	)
	cmd.Flags().StringVarP(&opts.Name, "name", "n", "", "explicit dependency name")

	return cmd
}
