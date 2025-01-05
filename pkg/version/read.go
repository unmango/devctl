package version

import (
	"errors"
	"fmt"
	"os"
	"strings"

	"github.com/spf13/afero"
	"github.com/unmango/devctl/pkg/version/internal"
	"github.com/unmango/devctl/pkg/version/opts"
)

func Read(name string, options ...opts.ReadOp) (string, error) {
	return ReadFile(RelPath(name), options...)
}

func ReadFile(path string, options ...opts.ReadOp) (string, error) {
	opts := internal.ReadOptions(options)

	if b, err := afero.ReadFile(opts.Fs, path); err == nil {
		return strings.TrimSpace(string(b)), nil
	} else if errors.Is(err, os.ErrNotExist) {
		return "", fmt.Errorf("dependency not found: %s", path)
	} else {
		return "", err
	}
}
