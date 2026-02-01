package internal

import (
	"github.com/unmango/devctl/pkg/version/opts"
	"github.com/unmango/go/fopt"
)

func InitOptions(options []opts.InitOp) *opts.Init {
	o := opts.DefaultInit
	fopt.ApplyAll(&o, options)
	return &o
}

func PrintOptions(options []opts.PrintOp) *opts.Print {
	o := opts.DefaultPrint
	fopt.ApplyAll(&o, options)
	return &o
}

func ReadOptions(options []opts.ReadOp) *opts.Read {
	o := opts.DefaultRead
	fopt.ApplyAll(&o, options)
	return &o
}

func WriteOptions(options []opts.WriteOp) *opts.Write {
	o := opts.DefaultWrite
	fopt.ApplyAll(&o, options)
	return &o
}
