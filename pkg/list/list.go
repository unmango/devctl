package list

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

func (o *Options) Printer(root string) *printer {
	return &printer{
		Opts:    o,
		Sources: o.sources(),
		Root:    root,
	}
}
