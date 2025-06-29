package main

import (
	"os"

	"github.com/charmbracelet/log"
	"github.com/unmango/devctl/cmd"
	"github.com/unmango/devctl/pkg/version"
	"github.com/unmango/go/cli"
)

func main() {
	log.SetLevel(log.ErrorLevel)

	if len(os.Args) == 2 {
		n, err := version.PrintIfPath(os.Args[1])
		if err != nil {
			cli.Fail(err)
		}
		if n > 0 {
			return
		}
	}

	if err := cmd.Execute(); err != nil {
		cli.Fail(err)
	}
}
