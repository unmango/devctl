package stack

import "context"

type AddOptions struct {
	CommitOptions
	Base string
}

// Add puts a new layer on top of the stack, creating the stack when there
// isn't one yet.
//
//	gh stack add [-A|-u] [-m <message>] <branch>
//
// Or, outside of a stack:
//
//	gh stack init [--base <base>] <branch>
//	git add -A|-u
//	git commit -m <message>
func Add(ctx context.Context, branch string, options AddOptions) error {
	stacked, err := InStack(ctx)
	if err != nil {
		return err
	}

	if stacked {
		return Stack(ctx, append([]string{"add"}, options.AddArgs(branch)...)...)
	}

	args := []string{"init"}
	if options.Base != "" {
		args = append(args, "--base", options.Base)
	}

	if err = Stack(ctx, append(args, branch)...); err != nil {
		return err
	}

	return options.Commit(ctx)
}
