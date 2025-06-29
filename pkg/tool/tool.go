package tool

import (
	"context"

	"github.com/unmango/devctl/pkg/work"
)

type Tool struct {
	Config

	Name string
}

func (t *Tool) Install(ctx context.Context, dir work.Directory) error {
	return nil
}
