package main

/*
#include <gnumake.h>
*/
import "C"

import (
	"unsafe"

	"github.com/unmango/devctl/pkg/make"
	"github.com/unmango/gnumake-go"
)

// https://www.gnu.org/software/make/manual/html_node/load-Directive.html

//export mk_funcs_gmk_setup
func mk_funcs_gmk_setup(floc *C.gmk_floc) int {
	gofloc := gnumake.NewFileLocationRef(unsafe.Pointer(floc))
	gofloc.Deref()

	make.RegisterFuncs(gofloc)
	return make.SetupSuccess
}

func main() {}
