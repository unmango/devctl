package internal

import (
	"github.com/unmango/devctl/pkg/version/opts"
	"github.com/unmango/go/option"
)

func InitOptions(options []opts.InitOp) *opts.Init {
	o := opts.DefaultInit
	option.ApplyAll(&o, options)
	return &o
}

func PrintOptions(options []opts.PrintOp) *opts.Print {
	o := opts.DefaultPrint
	option.ApplyAll(&o, options)
	return &o
}

func ReadOptions(options []opts.ReadOp) *opts.Read {
	o := opts.DefaultRead
	option.ApplyAll(&o, options)
	return &o
}

func WriteOptions(options []opts.WriteOp) *opts.Write {
	o := opts.DefaultWrite
	option.ApplyAll(&o, options)
	return &o
}
