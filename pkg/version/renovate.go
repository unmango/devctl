package version

import (
	"encoding/json"
	"os"
	"path/filepath"

	"github.com/spf13/afero"
	"github.com/unmango/devctl/pkg/renovate"
	"github.com/unmango/devctl/pkg/version/internal"
	"github.com/unmango/devctl/pkg/version/opts"
)

var RenovatePath = filepath.Join(DirName, "renovate.config")

func RenovateManager(name string) *renovate.CustomManager {
	return &renovate.CustomManager{
		CustomType:   "regex",
		FileMatch:    []string{filepath.Join(DirName, name)},
		MatchStrings: []string{"(?<currentValue>+)"},
	}
}

func RenovateConfig(name string) *renovate.Config {
	return &renovate.Config{CustomManagers: []interface{}{
		RenovateManager(name),
	}}
}

func WriteRenovateConfig(config *renovate.Config, options ...opts.WriteOp) error {
	opts := internal.WriteOptions(options)

	data, err := json.Marshal(config)
	if err != nil {
		return err
	}

	return afero.WriteFile(opts.Fs,
		RenovatePath,
		data,
		os.ModePerm,
	)
}
