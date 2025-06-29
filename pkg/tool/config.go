package tool

type Config struct {
	Url    string `json:"url,omitempty"`
	Script bool   `json:"script,omitempty"`
}

func FromConfig(name string, config Config) (*Tool, error) {
	return &Tool{Config: config, Name: name}, nil
}
