package stack

import (
	"github.com/spf13/cobra"
	"github.com/unmango/devctl/pkg/stack"
	"github.com/unmango/devctl/pkg/work"
	"github.com/unmango/go/cli"
)

var SaveCmd = NewSave()

type SaveOptions struct {
	stack.SaveOptions
	work.ChdirOptions
}

func NewSave() *cobra.Command {
	opts := SaveOptions{}

	cmd := &cobra.Command{
		Use: "save",
		Example: `devctl stack save -Am "Add the thing"
devctl stack save --no-push`,
		Short: "Commit the current work and propagate it through the stack",
		Long: `Commit the current work and propagate it through the stack.

	git add -A|-u
	git commit -m <message>
	gh stack rebase --upstack
	gh stack push

Staging and committing are skipped unless -A, -u, or -m are given.`,
		Aliases: []string{"s"},
		Args:    cobra.NoArgs,
		Run: func(cmd *cobra.Command, args []string) {
			ctx := cmd.Context()
			if err := opts.Chdir(ctx); err != nil {
				cli.Fail(err)
			}

			if err := stack.Save(ctx, opts.SaveOptions); err != nil {
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
