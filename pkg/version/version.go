package version

import (
	"fmt"
	"path/filepath"
	"regexp"
	"strings"
)

const DirName = ".versions"

var Regex = regexp.MustCompile(`v?\d+\.\d+\.\d+`)

func RelPath(name string) string {
	return filepath.Join(DirName, name)
}

func Clean(version string) string {
	return strings.TrimPrefix(version, "v")
}

func IsPath(name string) bool {
	return strings.HasPrefix(name, DirName)
}

func Prefixed(name string) string {
	return fmt.Sprint("v", name)
}
