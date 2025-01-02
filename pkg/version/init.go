package version

import (
	"context"
	"errors"
	"fmt"
	"os"
	"path/filepath"

	"github.com/spf13/afero"
	"github.com/unmango/go/option"
)

func Init(ctx context.Context, name string, src Source, options ...Option) (err error) {
	opts := Options{fs: afero.NewOsFs()}
	option.ApplyAll(&opts, options)

	if name == "" {
		if name, err = src.Name(ctx); err != nil {
			return fmt.Errorf("name is required")
		}
	}

	err = opts.fs.Mkdir(DirName, os.ModePerm)
	if err != nil && !errors.Is(err, os.ErrExist) {
		return
	}

	var version string
	if version, err = src.Latest(ctx); err != nil {
		return
	}

	return afero.WriteFile(opts.fs,
		filepath.Join(DirName, name),
		[]byte(Clean(version)+"\n"),
		os.ModePerm,
	)
}
