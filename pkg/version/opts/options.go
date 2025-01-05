package opts

import "github.com/spf13/afero"

type Init struct {
	Fs afero.Fs
}

var DefaultInit = Init{
	Fs: afero.NewOsFs(),
}

type InitOp func(*Init)

func WithInitFs(fs afero.Fs) InitOp {
	return func(o *Init) {
		o.Fs = fs
	}
}

type Print struct {
	Clean    bool
	NewLine  bool
	Prefixed bool
}

var DefaultPrint = Print{
	Clean:    true,
	NewLine:  true,
	Prefixed: true,
}

type PrintOp func(*Print)

func IncludePrefix(o *Print) {
	o.Prefixed = true
}

func IncludeNewLine(o *Print) {
	o.NewLine = true
}

func PrintPrefixed(prefixed bool) PrintOp {
	return func(o *Print) {
		o.Prefixed = prefixed
	}
}

func PrintClean(clean bool) PrintOp {
	return func(o *Print) {
		o.Clean = clean
	}
}

func PrintNewLine(newLine bool) PrintOp {
	return func(o *Print) {
		o.NewLine = newLine
	}
}

type Read struct {
	Fs afero.Fs
}

var DefaultRead = Read{
	Fs: afero.NewOsFs(),
}

type ReadOp func(*Read)

func WithReadFs(fs afero.Fs) ReadOp {
	return func(o *Read) {
		o.Fs = fs
	}
}

type Write struct {
	Fs afero.Fs
}

var DefaultWrite = Write{
	Fs: afero.NewOsFs(),
}

type WriteOp func(*Write)

func WithWriteFs(fs afero.Fs) WriteOp {
	return func(o *Write) {
		o.Fs = fs
	}
}
