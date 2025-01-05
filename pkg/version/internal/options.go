package internal

import (
	"github.com/unmango/devctl/pkg/version/opts"
	"github.com/unmango/go/option"
)

func PrintOptions(options []opts.PrintOp) *opts.Print {
	o := opts.DefaultPrint
	option.ApplyAll(&o, options)
	return &o
}

func InitOptions(options []opts.InitOp) *opts.Init {
	o := opts.DefaultInit
	option.ApplyAll(&o, options)
	return &o
}
