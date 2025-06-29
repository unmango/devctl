package config

import (
	"encoding/json"

	"github.com/spf13/afero"
	"github.com/unmango/devctl/pkg/work"
)

const (
	DefaultName = "devctl"
	DefaultExt  = ".yml"
	DefaultFile = DefaultName + DefaultExt
)

type Config struct{}

func (c *Config) Marshal() ([]byte, error) {
	return json.Marshal(c)
}

var Empty = &Config{}

type Options struct {
	fs afero.Fs
}

type Option func(*Options)

func Init(dir work.Directory) error {
	return Viper(dir).SafeWriteConfigAs(DefaultFile)
}

func Load(dir work.Directory) (*Config, error) {
	if err := Viper(dir).ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
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

func Viper(dir work.Directory) *viper.Viper {
	v := viper.New()
	v.SetFs(dir.Fs())
	v.SetConfigName("devctl")
	v.AddConfigPath(dir.Path())

	return v
}
