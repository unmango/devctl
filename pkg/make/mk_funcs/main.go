package main

import (
	"C"

	"github.com/unmango/devctl/pkg/make"
)

// https://www.gnu.org/software/make/manual/html_node/load-Directive.html

//export mk_funcs_gmk_setup
func mk_funcs_gmk_setup() int {
	make.RegisterFuncs()
	return make.SetupSuccess
}

func main() {}
