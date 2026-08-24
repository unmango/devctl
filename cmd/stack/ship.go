package stack

import (
	"github.com/spf13/cobra"
	"github.com/unmango/devctl/pkg/stack"
	"github.com/unmango/devctl/pkg/work"
	"github.com/unmango/go/cli"
)

var ShipCmd = NewShip()

type ShipOptions struct {
	stack.ShipOptions
	work.ChdirOptions
}

func NewShip() *cobra.Command {
	opts := ShipOptions{}

	cmd := &cobra.Command{
		Use: "ship",
		Example: `devctl stack ship
devctl stack ship --open --prune`,
		Short: "Sync the stack with the remote and open or update its pull requests",
		Long: `Sync the stack with the remote and open or update its pull requests.

	gh stack sync [--prune]
	gh stack submit --auto [--open]

New pull requests are created as drafts unless --open is given. Note that sync
succeeds while printing "Sync aborted" when the local and remote stacks have
diverged, which this command cannot detect.`,
		Aliases: []string{"submit"},
		Args:    cobra.NoArgs,
		Run: func(cmd *cobra.Command, args []string) {
			ctx := cmd.Context()
			if err := opts.Chdir(ctx); err != nil {
				cli.Fail(err)
			}

			if err := stack.Ship(ctx, opts.ShipOptions); err != nil {
				fail(err)
			}
		},
	}

	_ = work.ChdirFlag(cmd, &opts.ChdirOptions, "")
	remoteFlag(cmd, &opts.RemoteOptions)
	cmd.Flags().BoolVar(&opts.NoSync, "no-sync", false, "submit without syncing first")
	cmd.Flags().BoolVar(&opts.Open, "open", false, "mark new and existing pull requests ready for review")
	cmd.Flags().BoolVar(&opts.Prune, "prune", false, "delete local branches for merged pull requests")

	return cmd
}
