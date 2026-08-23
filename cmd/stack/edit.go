package stack

import (
	"github.com/spf13/cobra"
	"github.com/unmango/devctl/pkg/stack"
	"github.com/unmango/devctl/pkg/work"
	"github.com/unmango/go/cli"
)

var EditCmd = NewEdit()

type EditOptions struct {
	stack.EditOptions
	work.ChdirOptions
}

func NewEdit() *cobra.Command {
	opts := EditOptions{}

	cmd := &cobra.Command{
		Use:     "edit [branch]",
		Example: `devctl stack edit feat/auth -Am "Fix the middleware"`,
		Short:   "Commit work to a lower layer without losing your place",
		Long: `Commit work to a lower layer without losing your place.

	gh stack checkout <branch>
	git add -A|-u
	git commit -m <message>
	gh stack rebase --upstack
	gh stack checkout <the branch you were on>
	gh stack push

Uncommitted changes travel with the checkout, so it makes no difference whether
they were staged before running this.`,
		Aliases: []string{"e"},
		Args:    cobra.ExactArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			ctx := cmd.Context()
			if err := opts.Chdir(ctx); err != nil {
				cli.Fail(err)
			}

			if err := stack.Edit(ctx, args[0], opts.EditOptions); err != nil {
				fail(err)
			}
		},
	}

	_ = work.ChdirFlag(cmd, &opts.ChdirOptions, "")
	commitFlags(cmd, &opts.CommitOptions)
	remoteFlag(cmd, &opts.RemoteOptions)
	cmd.Flags().BoolVar(&opts.NoPush, "no-push", false, "skip pushing the stack")

	return cmd
}
