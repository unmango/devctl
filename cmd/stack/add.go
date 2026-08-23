package stack

import (
	"github.com/spf13/cobra"
	"github.com/unmango/devctl/pkg/stack"
	"github.com/unmango/devctl/pkg/work"
	"github.com/unmango/go/cli"
)

var AddCmd = NewAdd()

type AddOptions struct {
	stack.AddOptions
	work.ChdirOptions
}

func NewAdd() *cobra.Command {
	opts := AddOptions{}

	cmd := &cobra.Command{
		Use: "add [branch]",
		Example: `devctl stack add feat/auth -Am "Add auth middleware"
devctl stack add feat/auth --base develop`,
		Short: "Put a new layer on top of the stack, creating the stack if needed",
		Long: `Put a new layer on top of the stack, creating the stack if needed.

	gh stack add [-A|-u] [-m <message>] <branch>

Or, when the current branch isn't part of a stack:

	gh stack init [--base <base>] <branch>
	git add -A|-u
	git commit -m <message>`,
		Aliases: []string{"new", "a"},
		Args:    cobra.ExactArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			ctx := cmd.Context()
			if err := opts.Chdir(ctx); err != nil {
				cli.Fail(err)
			}

			if err := stack.Add(ctx, args[0], opts.AddOptions); err != nil {
				fail(err)
			}
		},
	}

	_ = work.ChdirFlag(cmd, &opts.ChdirOptions, "")
	commitFlags(cmd, &opts.CommitOptions)
	cmd.Flags().StringVarP(&opts.Base, "base", "b", "",
		"trunk branch, only used when creating the stack",
	)

	return cmd
}
