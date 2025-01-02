package version

import (
	"errors"
	"fmt"
	"os"
	"strings"

	"github.com/spf13/afero"
	"github.com/unmango/go/option"
)

func ReadFile(name string, options ...Option) (string, error) {
	opts := &Options{fs: afero.NewOsFs()}
	option.ApplyAll(opts, options)

	if b, err := afero.ReadFile(opts.fs, RelPath(name)); err == nil {
		return strings.TrimSpace(string(b)), nil
	} else if errors.Is(err, os.ErrNotExist) {
		return "", fmt.Errorf("dependency not found: %s", name)
	} else {
		return "", err
	}
}
