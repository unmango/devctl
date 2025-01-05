package version

import (
	"context"
	"errors"
	"fmt"
	"os"
	"path/filepath"

	"github.com/spf13/afero"
	"github.com/unmango/devctl/pkg/version/internal"
	"github.com/unmango/devctl/pkg/version/opts"
)

func Init(ctx context.Context, name string, src Source, options ...opts.InitOp) (err error) {
	opts := internal.InitOptions(options)

	if name == "" {
		if name, err = src.Name(ctx); err != nil {
			return fmt.Errorf("name is required")
		}
	}

	err = opts.Fs.Mkdir(DirName, os.ModePerm)
	if err != nil && !errors.Is(err, os.ErrExist) {
		return
	}

	var version string
	if version, err = src.Latest(ctx); err != nil {
		return
	}

	return afero.WriteFile(opts.Fs,
		filepath.Join(DirName, name),
		[]byte(Clean(version)+"\n"),
		os.ModePerm,
	)
}
