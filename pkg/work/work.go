package work

import (
	"context"
	"fmt"
	"os"

	"github.com/charmbracelet/log"
	"github.com/unmango/go/vcs/git"
)

type cwdKey struct{}

// A Directory defines a path to a valid, existing directory on the local filesystem
type Directory string

// Path returns the [Directory] path as a string
func (c Directory) Path() string {
	return string(c)
}

// Git returns a [Directory] pointing to the git repository closest to the current working directory
func Git(ctx context.Context) (work Directory, err error) {
	if p, err := git.Root(ctx); err != nil {
		return "", err
	} else {
		return Directory(p), nil
	}
}

// Cwd returns a [Directory] pointing to the current working directory
func Cwd() (work Directory, err error) {
	if p, err := os.Getwd(); err != nil {
		return "", err
	} else {
		return Directory(p), nil
	}
}

// Load returns a [Directory] pointing to the first directory able to be resolved without error.
//
// Load will attempt directories in the following order:
//   - [Git]
//   - [Cwd]
func Load(ctx context.Context) (work Directory, err error) {
	if work, err = Git(ctx); err == nil {
		return
	} else {
		log.Debugf("loading git context: %s", err)
	}

	if work, err = Cwd(); err == nil {
		return
	} else {
		log.Debugf("loading cwd context: %s", err)
	}

	return "", fmt.Errorf("failed to load current work context")
}
