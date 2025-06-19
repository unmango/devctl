package make

/*
#include <stdlib.h>
#include <string.h>
*/
import "C"

import (
	"fmt"

	"github.com/unmango/gnumake-go"
)

var fileLocation *gnumake.FileLocation

func Test(nm string, argc uint32, argv []*byte) []byte {
	fmt.Println("nm", nm)
	fmt.Println("argc", argc)
	fmt.Println("argv", argv)
	return nil
}

func RegisterFuncs(floc *gnumake.FileLocation) {
	fmt.Println("floc", floc.LineNumber)
	gnumake.AddFunction("testfunc", Test, 1, 2, 0)
}
