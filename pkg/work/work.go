package work

import (
	"context"
	"fmt"
	"os"

	"github.com/charmbracelet/log"
	"github.com/spf13/afero"
	"github.com/unmango/go/option"
	"github.com/unmango/go/vcs/git"
)

// A Directory defines a path to a valid, existing directory on the local filesystem
type Directory string

type Options struct {
	fs afero.Fs
}

type Option func(*Options)

// WithRoot returns an option to use fs as the root for filesystem operations
func WithRoot(fs afero.Fs) Option {
	return func(o *Options) {
		o.fs = fs
	}
}

// Path returns the [Directory] path as a string
func (c Directory) Path() string {
	return string(c)
}

// Fs returns an [afero.Fs] rooted at c
func (c Directory) Fs(options ...Option) afero.Fs {
	opts := &Options{fs: afero.NewOsFs()}
	option.ApplyAll(opts, options)

	return afero.NewBasePathFs(opts.fs, c.Path())
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
