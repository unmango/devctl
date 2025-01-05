package version

import (
	"fmt"
	"path/filepath"
	"regexp"
	"strings"

	"github.com/spf13/afero"
)

const DirName = ".versions"

var Regex = regexp.MustCompile(`v?\d+\.\d+\.\d+`)

var DefaultOptions = Options{
	fs:       afero.NewOsFs(),
	prefixed: true,
}

type Options struct {
	fs       afero.Fs
	prefixed bool
}

type Option func(*Options)

func WithFs(fs afero.Fs) Option {
	return func(o *Options) {
		o.fs = fs
	}
}

func IncludePrefix(opts *Options) {
	opts.prefixed = true
}

func WithPrefixed(prefixed bool) Option {
	return func(o *Options) {
		o.prefixed = prefixed
	}
}

func RelPath(name string) string {
	return filepath.Join(DirName, name)
}

func Clean(version string) string {
	return strings.TrimPrefix(version, "v")
}

func IsPath(name string) bool {
	return strings.HasPrefix(name, DirName)
}

func Prefixed(name string) string {
	return fmt.Sprint("v", name)
}
