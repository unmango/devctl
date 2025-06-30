package tool

import "fmt"

type Config struct {
	Script  bool   `json:"script,omitempty"`
	Url     string `json:"url,omitempty"`
	Version string `json:"version,omitempty"`
}

func (c Config) Verify() error {
	if c.Url == "" {
		return fmt.Errorf("no url")
	}

	return nil
}

func FromConfig(name string, config Config) (*Tool, error) {
	if err := config.Verify(); err != nil {
		return nil, err
	}

	return &Tool{
		Config: config,
		Name:   name,
	}, nil
}
