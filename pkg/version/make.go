package version

import (
	"fmt"
	"os"
	"strings"

	"github.com/spf13/afero"
	"github.com/unmango/go/option"
)

func MakefileVar(name string) string {
	return fmt.Sprintf("%s_VERSION := $(shell devctl version %s)\n",
		strings.ToUpper(name), name,
	)
}

func WriteMakefile(name string, options ...Option) error {
	opts := &Options{fs: afero.NewOsFs()}
	option.ApplyAll(opts, options)

	return afero.WriteFile(opts.fs,
		fmt.Sprintf("%s/%s.mk", DirName, name),
		[]byte(MakefileVar(name)),
		os.ModePerm,
	)
}
