package list

import (
	"fmt"
	"path/filepath"
	"slices"
	"strings"
)

type printer struct {
	Opts    *Options
	Sources []string
	Root    string
}

func (p *printer) shouldPrint(path string) bool {
	// TODO: No sources provided && exclude tests
	if len(p.Sources) == 0 {
		return true
	}

	ext := filepath.Ext(path)
	if !slices.Contains(p.Sources, ext) {
		return false
	}

	switch ext {
	case ".go":
		return p.shouldPrintGo(path)
	case ".ts":
		return p.shouldPrintTs(path)
	case ".cs":
		return p.shouldPrintDotnet(path)
	case ".fs":
		return p.shouldPrintDotnet(path)
	}

	return true
}

func (p *printer) shouldPrintGo(path string) bool {
	if p.Opts.ExcludeTests {
		return !strings.Contains(path, "_test.go")
	}

	return true
}

func (p *printer) shouldPrintTs(path string) bool {
	if p.Opts.ExcludeTests {
		return !strings.Contains(path, ".spec.ts")
	}

	return true
}

func (p *printer) shouldPrintDotnet(path string) bool {
	if strings.Contains(path, "/bin/") || strings.Contains(path, "/obj/") {
		return false
	}
	if p.Opts.ExcludeTests {
		matched, err := filepath.Match("**/*.Tests?.*", path)
		if err != nil {
			panic(err)
		}

		return !matched
	}

	return true
}

func (p *printer) handle(path string) (err error) {
	if !p.shouldPrint(path) {
		return nil
	}

	if !p.Opts.Absolute {
		path, err = filepath.Rel(p.Root, path)
	}
	if err != nil {
		return err
	}

	fmt.Println(path)
	return nil
}
