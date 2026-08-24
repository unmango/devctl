package stack

import "context"

type SaveOptions struct {
	CommitOptions
	RemoteOptions
	NoPush bool
}

// Save commits the current work and propagates it through the stack.
//
//	git add -A|-u
//	git commit -m <message>
//	gh stack rebase --upstack
//	gh stack push
func Save(ctx context.Context, options SaveOptions) error {
	if err := options.Commit(ctx); err != nil {
		return err
	}

	if err := Stack(ctx, options.Args("rebase", "--upstack")...); err != nil {
		return err
	}

	if options.NoPush {
		return nil
	}

	return Stack(ctx, options.Args("push")...)
}
