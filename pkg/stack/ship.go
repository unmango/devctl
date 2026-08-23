package stack

import "context"

type ShipOptions struct {
	RemoteOptions
	NoSync bool
	Open   bool
	Prune  bool
}

// Ship syncs the stack with the remote and opens or updates its pull requests.
//
//	gh stack sync [--prune]
//	gh stack submit --auto [--open]
//
// submit always runs with --auto because a bare submit opens a full screen
// editor when attached to a terminal.
func Ship(ctx context.Context, options ShipOptions) error {
	if !options.NoSync {
		args := []string{"sync"}
		if options.Prune {
			args = append(args, "--prune")
		}

		if err := Stack(ctx, options.Args(args...)...); err != nil {
			return err
		}
	}

	args := []string{"submit", "--auto"}
	if options.Open {
		args = append(args, "--open")
	}

	return Stack(ctx, options.Args(args...)...)
}
