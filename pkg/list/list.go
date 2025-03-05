package list

import (
	"io/fs"
	"path/filepath"
	"slices"
	"strings"
)

var Blacklist = []string{
	"node_modules",
	"bin", "obj",
	"pcl",
	".tdl-old",
	".uml2ts-old",
	"testdata",
	".idea",
	".vscode",
	".git",
}

type Options struct {
	Absolute     bool
	ExcludeTests bool
	Go           bool
	Proto        bool
	Typescript   bool
	CSharp       bool
	FSharp       bool
	Dotnet       bool
}

func (o *Options) sources() []string {
	sources := []string{}
	if o.Go {
		sources = append(sources, ".go")
	}
	if o.Proto {
		sources = append(sources, ".proto")
	}
	if o.Typescript {
		sources = append(sources, ".ts")
	}
	if o.Dotnet || o.CSharp {
		sources = append(sources, ".cs")
	}
	if o.Dotnet || o.FSharp {
		sources = append(sources, ".fs")
	}

	return sources
}

func (o *Options) printer(root string) *printer {
	return &printer{
		Opts:    o,
		Sources: o.sources(),
		Root:    root,
	}
}

func Directory(root string, options *Options) error {
	printer := options.printer(root)
	return filepath.WalkDir(root,
		func(path string, d fs.DirEntry, err error) error {
			if d.IsDir() {
				if blacklisted(path) {
					return filepath.SkipDir
				}

				return nil
			}

			return printer.handle(path)
		},
	)
}

func blacklisted(path string) bool {
	return slices.ContainsFunc(Blacklist, func(b string) bool {
		return strings.Contains(path, b)
	})
}
