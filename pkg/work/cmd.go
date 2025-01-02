package work

import (
	"context"
	"os"

	"github.com/spf13/cobra"
)

type ChdirOptions struct {
	dir string
}

func (o ChdirOptions) Cwd(ctx context.Context) (Directory, error) {
	if o.dir != "" {
		return Directory(o.dir), nil
	} else {
		return Load(ctx)
	}
}

func (o ChdirOptions) Chdir(ctx context.Context) error {
	if d, err := o.Cwd(ctx); err != nil {
		return err
	} else {
		return os.Chdir(d.Path())
	}
}

func NewChdirOptions(d string) ChdirOptions {
	return ChdirOptions{dir: d}
}

func ChdirFlag(cmd *cobra.Command, opts *ChdirOptions, value string) error {
	cmd.Flags().StringVarP(&opts.dir, "chdir", "C", value,
		"change to the specified directory before executing",
	)
	return cmd.MarkFlagDirname("chdir")
}
