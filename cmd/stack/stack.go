package stack

import (
	"os"

	"github.com/spf13/cobra"
	"github.com/unmango/devctl/pkg/stack"
	"github.com/unmango/go/cli"
)

var Cmd = New()

func init() {
	Cmd.AddCommand(
		AddCmd,
		EditCmd,
		SaveCmd,
		ShipCmd,
	)
}

func New() *cobra.Command {
	return &cobra.Command{
		Use:   "stack",
		Short: "Orchestrate gh-stack commands for working with stacked diffs",
		Long: `Orchestrate gh-stack commands for working with stacked diffs.

Each subcommand runs a fixed sequence of git and gh stack commands, nothing
more. Requires the gh-stack extension:

	gh extension install github/gh-stack`,
		Aliases: []string{"st"},
	}
}

// fail exits with the same code as the gh or git process that failed.
//
// gh stack exit codes are meaningful (2: not in a stack, 3: rebase conflict,
// 7: rebase in progress) and the failed command has already explained itself
// on stderr.
func fail(err error) {
	if code := stack.ExitCode(err); code > 0 {
		os.Exit(code)
	}

	cli.Fail(err)
}

func commitFlags(cmd *cobra.Command, options *stack.CommitOptions) {
	cmd.Flags().BoolVarP(&options.All, "all", "A", false,
		"stage all changes including untracked files",
	)
	cmd.Flags().BoolVarP(&options.Update, "update", "u", false,
		"stage changes to tracked files only",
	)
	cmd.Flags().StringVarP(&options.Message, "message", "m", "",
		"create a commit with this message",
	)
	cmd.MarkFlagsMutuallyExclusive("all", "update")
}

func remoteFlag(cmd *cobra.Command, options *stack.RemoteOptions) {
	cmd.Flags().StringVar(&options.Remote, "remote", "",
		"remote to operate against, required when the repo has multiple remotes",
	)
}
