package migrate

import (
	"github.com/spf13/cobra"
	"github.com/unmango/devctl/pkg/agents"
	"github.com/unmango/devctl/pkg/work"
	"github.com/unmango/go/cli"
)

var AgentsCmd = NewAgents()

type AgentsOptions struct {
	agents.Options
	work.ChdirOptions
}

func NewAgents() *cobra.Command {
	opts := AgentsOptions{}

	cmd := &cobra.Command{
		Use: "agents",
		Example: `devctl migrate agents
devctl migrate agents -C ../other-repo --force`,
		Short: "Migrate a CLAUDE.md to the AGENTS.md convention",
		Long: `Migrate a CLAUDE.md to the AGENTS.md convention.

The instructions ` + "`claude /init`" + ` writes are useful to every coding agent, only
the filename and a few boilerplate phrases are Claude specific. Three files
come out of the migration:

	AGENTS.md                        the instructions, boilerplate rewritten
	CLAUDE.md                        imports AGENTS.md
	.github/copilot-instructions.md  points at AGENTS.md

Lines still mentioning Claude after the rewrite are reported rather than
guessed at, edit those by hand.`,
		Aliases: []string{"agent"},
		Args:    cobra.NoArgs,
		Run: func(cmd *cobra.Command, args []string) {
			dir, err := opts.Cwd(cmd.Context())
			if err != nil {
				cli.Fail(err)
			}

			remaining, err := agents.Migrate(dir.Fs(), opts.Options)
			if err != nil {
				cli.Fail(err)
			}

			for _, ref := range remaining {
				cmd.PrintErrf("%s:%d still mentions Claude: %s\n",
					agents.AgentsFile, ref.Line, ref.Text,
				)
			}
		},
	}

	_ = work.ChdirFlag(cmd, &opts.ChdirOptions, "")
	cmd.Flags().BoolVar(&opts.Force, "force", false,
		"overwrite an existing AGENTS.md and copilot instructions",
	)

	return cmd
}
