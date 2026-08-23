// Package stack orchestrates sequences of gh-stack commands.
//
// Nothing here reimplements gh stack. Every exported function is a fixed list
// of git and gh stack invocations with stdio wired straight through, so the
// underlying tools own all output, prompting, and exit codes.
package stack

import (
	"context"
	"errors"
	"io"
	"os"
	"os/exec"
	"strings"

	"github.com/charmbracelet/log"
)

type (
	ghKey  struct{}
	gitKey struct{}
)

// NotInStackCode is the code gh stack exits with outside of a stack
const NotInStackCode = 2

// WithGhPath returns a context resolving the gh binary to path
func WithGhPath(parent context.Context, path string) context.Context {
	return context.WithValue(parent, ghKey{}, path)
}

// WithGitPath returns a context resolving the git binary to path
func WithGitPath(parent context.Context, path string) context.Context {
	return context.WithValue(parent, gitKey{}, path)
}

// Stack runs `gh stack` with args
func Stack(ctx context.Context, args ...string) error {
	return Gh(ctx, append([]string{"stack"}, args...)...)
}

// Gh runs gh with args
func Gh(ctx context.Context, args ...string) error {
	gh, err := ghPath(ctx)
	if err != nil {
		return err
	}

	return run(exec.CommandContext(ctx, gh, args...))
}

// Git runs git with args
func Git(ctx context.Context, args ...string) error {
	git, err := gitPath(ctx)
	if err != nil {
		return err
	}

	return run(exec.CommandContext(ctx, git, args...))
}

// GitOutput runs git with args and returns its trimmed standard output
func GitOutput(ctx context.Context, args ...string) (string, error) {
	git, err := gitPath(ctx)
	if err != nil {
		return "", err
	}

	cmd := exec.CommandContext(ctx, git, args...)
	log.Debug(strings.Join(cmd.Args, " "))
	cmd.Stderr = os.Stderr

	out, err := cmd.Output()
	return strings.TrimSpace(string(out)), err
}

// CurrentBranch returns the name of the currently checked out branch
func CurrentBranch(ctx context.Context) (string, error) {
	return GitOutput(ctx, "rev-parse", "--abbrev-ref", "HEAD")
}

// InStack reports whether the current branch belongs to a stack
func InStack(ctx context.Context) (bool, error) {
	gh, err := ghPath(ctx)
	if err != nil {
		return false, err
	}

	// view writes the stack to stdout, we only care about the exit code
	cmd := exec.CommandContext(ctx, gh, "stack", "view", "--json")
	log.Debug(strings.Join(cmd.Args, " "))
	cmd.Stdout = io.Discard
	cmd.Stderr = io.Discard

	if err = cmd.Run(); err == nil {
		return true, nil
	} else if ExitCode(err) == NotInStackCode {
		return false, nil
	} else {
		return false, err
	}
}

// ExitCode returns the exit code of the process err came from, or -1 when err
// was not produced by a child process
func ExitCode(err error) int {
	var exit *exec.ExitError
	if errors.As(err, &exit) {
		return exit.ExitCode()
	}

	return -1
}

func run(cmd *exec.Cmd) error {
	log.Debug(strings.Join(cmd.Args, " "))
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	return cmd.Run()
}

func ghPath(ctx context.Context) (string, error) {
	if gh, ok := ctx.Value(ghKey{}).(string); ok && gh != "" {
		return gh, nil
	}

	if gh := os.Getenv("GH_PATH"); gh != "" {
		return gh, nil
	}

	if gh, err := exec.LookPath("gh"); err != nil {
		return "", errors.New("stacked diffs require the GitHub CLI: https://cli.github.com")
	} else {
		return gh, nil
	}
}

func gitPath(ctx context.Context) (string, error) {
	if git, ok := ctx.Value(gitKey{}).(string); ok && git != "" {
		return git, nil
	}

	if git := os.Getenv("GIT_PATH"); git != "" {
		return git, nil
	}

	return exec.LookPath("git")
}
