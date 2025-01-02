package version

import (
	"path/filepath"
	"regexp"
	"strings"

	"github.com/spf13/afero"
)

const DirName = ".versions"

var Regex = regexp.MustCompile(`v?\d+\.\d+\.\d+`)

type Options struct {
	fs afero.Fs
}

type Option func(*Options)

func WithFs(fs afero.Fs) Option {
	return func(o *Options) {
		o.fs = fs
	}
}

func RelPath(name string) string {
	return filepath.Join(DirName, name)
}

func Clean(version string) string {
	return strings.TrimPrefix(version, "v")
}
