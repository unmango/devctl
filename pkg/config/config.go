package config

import (
	"encoding/json"

	"github.com/spf13/afero"
	"github.com/unmango/devctl/pkg/work"
)

var DefaultName = ".devctl.json"

type Config struct{}

func (c *Config) Marshal() ([]byte, error) {
	return json.Marshal(c)
}

var Empty = &Config{}

type Options struct {
	fs afero.Fs
}

type Option func(*Options)

func (o *Options) workOpts() work.Option {
	if o.fs != nil {
		return work.WithRoot(o.fs)
	} else {
		return func(o *work.Options) {}
	}
}
