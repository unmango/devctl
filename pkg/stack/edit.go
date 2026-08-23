package stack

import "context"

type EditOptions struct {
	CommitOptions
	RemoteOptions
	NoPush bool
}

// Edit commits work to a lower layer of the stack and returns to the branch it
// started on.
//
//	gh stack checkout <branch>
//	git add -A|-u
//	git commit -m <message>
//	gh stack rebase --upstack
//	gh stack checkout <original branch>
//	gh stack push
//
// Uncommitted changes travel with the checkout, so staging them beforehand is
// equivalent to staging them here.
func Edit(ctx context.Context, branch string, options EditOptions) error {
	current, err := CurrentBranch(ctx)
	if err != nil {
		return err
	}

	if current != branch {
		if err = Stack(ctx, "checkout", branch); err != nil {
			return err
		}
	}

	if err = options.Commit(ctx); err != nil {
		return err
	}

	if err = Stack(ctx, options.Args("rebase", "--upstack")...); err != nil {
		return err
	}

	if current != branch {
		if err = Stack(ctx, "checkout", current); err != nil {
			return err
		}
	}

	if options.NoPush {
		return nil
	}

	return Stack(ctx, options.Args("push")...)
}
