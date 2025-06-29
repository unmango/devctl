package config

import (
	"log/slog"

	"github.com/charmbracelet/log"
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

var Empty = &Config{}

type Config struct {
	Tools map[string]tool.Config `json:"tools,omitempty"`
}

type Options struct {
	fs afero.Fs
}

type Option func(*Options)

func FromDirectory(dir work.Directory) (*Config, error) {
	return Unmarshal(Viper(dir))
}

func Init(dir work.Directory) error {
	return Viper(dir).SafeWriteConfigAs(DefaultFile)
}

func Unmarshal(viper *viper.Viper) (*Config, error) {
	if err := viper.ReadInConfig(); err != nil {
		return nil, err
	}

	var config Config
	if err := viper.Unmarshal(&config); err != nil {
		return nil, err
	} else {
		return &config, nil
	}
}

func Viper(dir work.Directory) *viper.Viper {
	v := viper.NewWithOptions(
		viper.WithLogger(slog.New(log.Default())),
	)

	// TODO: Connect this to my janky `work` package
	// v.SetFs(dir.Fs())
	v.SetConfigName("devctl")
	v.AddConfigPath(dir.Path())
	v.AddConfigPath(".")

	return v
}
