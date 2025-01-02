package version

import (
	"context"
	"fmt"
	"strings"

	"github.com/unmango/go/fs/github"
	"github.com/unmango/go/lazy"
)

type Source interface {
	Latest(context.Context) (string, error)
	Name(context.Context) (string, error)
}

func GuessSource(s string) (Source, error) {
	switch {
	case Regex.MatchString(s):
		return String(s), nil
	case strings.Contains(s, "github"):
		return GitHub(s), nil
	default:
		return nil, fmt.Errorf("unrecognized source: %s", s)
	}
}

type GitHub string

var ghclient = lazy.Once(func() *github.Client {
	return github.NewClient(nil)
})

// Name implements Source.
func (g GitHub) Name(context.Context) (string, error) {
	panic("unimplemented")
}

// Latest implements Source.
func (g GitHub) Latest(context.Context) (string, error) {
	panic("unimplemented")
}

type String string

// Latest implements Source.
func (s String) Latest(context.Context) (string, error) {
	return string(s), nil
}

// Name implements Source.
func (s String) Name(context.Context) (string, error) {
	return "", fmt.Errorf("invalid name: %s", s)
}
