package renovate

import "encoding/json"

func Unmarshal(data []byte) (*Config, error) {
	var config Config
	if err := json.Unmarshal(data, &config); err != nil {
		return nil, err
	} else {
		return &config, nil
	}
}

func Marshal(c *Config) ([]byte, error) {
	return json.Marshal(c)
}
