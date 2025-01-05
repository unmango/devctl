package version

import (
	"fmt"
	"os"
	"strings"

	"github.com/spf13/afero"
	"github.com/unmango/devctl/pkg/version/internal"
	"github.com/unmango/devctl/pkg/version/opts"
)

func MakefileVar(name string) string {
	return fmt.Sprintf("%s_VERSION := $(shell devctl version %s)\n",
		strings.ToUpper(name), name,
	)
}

func WriteMakefile(name string, options ...opts.WriteOp) error {
	opts := internal.WriteOptions(options)

	return afero.WriteFile(opts.Fs,
		fmt.Sprintf("%s/%s.mk", DirName, name),
		[]byte(MakefileVar(name)),
		os.ModePerm,
	)
}
