package config

import (
	"github.com/spf13/afero"
	"github.com/spf13/viper"
	"github.com/unmango/devctl/pkg/tool"
	"github.com/unmango/devctl/pkg/work"
)

const (
	DefaultName = "devctl"
	DefaultExt  = ".yml"
	DefaultFile = DefaultName + DefaultExt
)

type NotFoundError = viper.ConfigFileNotFoundError
type Config struct {
	Tools  map[string]tool.Config `json:"tools,omitempty"`
}

var Empty = &Config{}

type Options struct {
	fs afero.Fs
}

type Option func(*Options)

func Init(viper *viper.Viper) error {
	return viper.SafeWriteConfigAs(DefaultFile)
}

func Load(viper *viper.Viper) (*Config, error) {
	if err := viper.ReadInConfig(); err != nil {
		if _, ok := err.(NotFoundError); ok {
			return Empty, nil
		} else {
			return nil, err
		}
	}

	var config Config
	if err := viper.Unmarshal(&config); err != nil {
		return nil, err
	} else {
		return &config, nil
	}
}

func Unmarshal(viper *viper.Viper) (cfg Config, err error) {
	if err := viper.Unmarshal(&cfg); err != nil {
		return cfg, err
	} else {
		return cfg, nil
	}
}

func Viper(dir work.Directory) *viper.Viper {
	v := viper.New()
	v.SetFs(dir.Fs())
	v.SetConfigName("devctl")
	v.AddConfigPath(dir.Path())

	return v
}
