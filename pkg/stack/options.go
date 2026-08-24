package stack

import "context"

// CommitOptions describes how work in the tree becomes a commit.
//
// The flags mirror `gh stack add`: All stages everything including untracked
// files, Update stages tracked files only.
type CommitOptions struct {
	All     bool
	Message string
	Update  bool
}

// RemoteOptions describes the remote gh stack operates against
type RemoteOptions struct {
	Remote string
}

// Commit stages the tree when All or Update is set and creates a commit when
// Message is non-empty. A commit is never created without a message.
func (o CommitOptions) Commit(ctx context.Context) error {
	if args := o.stage(); args != nil {
		if err := Git(ctx, args...); err != nil {
			return err
		}
	}

	if o.Message == "" {
		return nil
	}

	return Git(ctx, "commit", "-m", o.Message)
}

// AddArgs returns the `gh stack add` flags equivalent to o, followed by branch
func (o CommitOptions) AddArgs(branch string) (args []string) {
	if o.All {
		args = append(args, "-A")
	} else if o.Update {
		args = append(args, "-u")
	}
	if o.Message != "" {
		args = append(args, "-m", o.Message)
	}

	return append(args, branch)
}

func (o CommitOptions) stage() []string {
	switch {
	case o.All:
		return []string{"add", "-A"}
	case o.Update:
		return []string{"add", "-u"}
	default:
		return nil
	}
}

// Args returns args with the remote appended when one was specified
func (o RemoteOptions) Args(args ...string) []string {
	if o.Remote == "" {
		return args
	}

	return append(args, "--remote", o.Remote)
}
