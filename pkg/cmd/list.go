package cmd

import (
	"fmt"
	"os"

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
			root, err := git.Root(cmd.Context())
			if err != nil {
				fmt.Fprintln(os.Stderr, err.Error())
				os.Exit(1)
			}

			err = list.Directory(root, options)
			if err != nil {
				fmt.Fprintln(os.Stderr, err)
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
