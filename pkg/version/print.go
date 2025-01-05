package version

import (
	"fmt"
	"io"
	"os"

	"github.com/unmango/devctl/pkg/version/internal"
	"github.com/unmango/devctl/pkg/version/opts"
)

func Cat(path string, w io.Writer, options ...opts.PrintOp) (int, error) {
	if v, err := ReadFile(path); err != nil {
		return 0, err
	} else {
		return Fprint(w, v, options...)
	}
}

func PrintIfPath(path string) (int, error) {
	if !IsPath(path) {
		return 0, nil
	}

	return Cat(path, os.Stdout,
		opts.IncludePrefix,
		opts.PrintNewLine(false),
	)
}

func Print(version string, options ...opts.PrintOp) (int, error) {
	return Fprint(os.Stdout, version, options...)
}

func Println(version string, options ...opts.PrintOp) (int, error) {
	return Fprintln(os.Stdout, version, options...)
}

func Fprint(w io.Writer, version string, options ...opts.PrintOp) (int, error) {
	opts := internal.PrintOptions(options)

	if opts.Clean {
		version = Clean(version)
	}
	if opts.Prefixed {
		version = Prefixed(version)
	}

	if opts.NewLine {
		return fmt.Fprintln(w, version)
	} else {
		return fmt.Fprint(w, version)
	}
}

func Fprintln(w io.Writer, version string, options ...opts.PrintOp) (int, error) {
	return Fprint(w, version, append(options, opts.IncludeNewLine)...)
}
