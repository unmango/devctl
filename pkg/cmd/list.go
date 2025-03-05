package cmd

import (
	"fmt"
	"io/fs"
	"os"
	"path/filepath"
	"slices"
	"strings"

	"github.com/charmbracelet/log"
	"github.com/spf13/cobra"
	"github.com/unmango/devctl/pkg/list"
	"github.com/unmango/go/vcs/git"
)

func NewList(options *list.Options) *cobra.Command {
	cmd := &cobra.Command{
		Use:     "list",
		Short:   "List source files in the current git repo",
		Aliases: []string{"ls"},
		Run: func(cmd *cobra.Command, args []string) {
			ctx := cmd.Context()
			log.Debug("running with options", "options", options)

			root, err := git.Root(ctx)
			if err != nil {
				fmt.Fprintln(os.Stderr, err.Error())
				os.Exit(1)
			}

			log.Debugf("walking root: %s", root)

			printer := options.Printer(root)
			err = filepath.WalkDir(root,
				func(path string, d fs.DirEntry, err error) error {
					if d.IsDir() {
						if blacklisted(path) {
							return filepath.SkipDir
						}

						return nil
					}

					return printer.Handle(path)
				},
			)
			if err != nil {
				fmt.Fprintln(os.Stderr, err.Error())
				os.Exit(1)
			}
		},
	}

	// TODO: It would probably make a lot more sense to have a e.g. --ext '.go' option
	cmd.Flags().BoolVar(&options.Absolute, "absolute", false, "Print fully qualified paths rather than paths relative to the git root")
	cmd.Flags().BoolVar(&options.ExcludeTests, "exclude-tests", false, "Exclude test files like *_test.go and *.spec.ts etc")
	cmd.Flags().BoolVar(&options.Go, "go", false, "List Go sources")
	cmd.Flags().BoolVar(&options.Typescript, "ts", false, "List TypeScript sources")
	cmd.Flags().BoolVar(&options.Proto, "proto", false, "List protobuf sources")
	cmd.Flags().BoolVar(&options.CSharp, "cs", false, "List C# sources")
	cmd.Flags().BoolVar(&options.FSharp, "fs", false, "List F# sources")
	cmd.Flags().BoolVar(&options.Dotnet, "dotnet", false, "List .NET sources")

	return cmd
}

func blacklisted(path string) bool {
	return slices.ContainsFunc(list.Blacklist, func(b string) bool {
		return strings.Contains(path, b)
	})
}
