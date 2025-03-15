package config

import (
	"os"

	"github.com/spf13/afero"
	"github.com/unmango/devctl/pkg/work"
	"github.com/unmango/go/option"
)

func Init(dir work.Directory, options ...Option) error {
	opts := &Options{}
	option.ApplyAll(opts, options)

	if config, err := Empty.Marshal(); err != nil {
		return err
	} else {
		return afero.WriteFile(
			dir.Fs(opts.workOpts()),
			DefaultName,
			config,
			os.ModePerm,
		)
	}
}
