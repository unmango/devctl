package make

/*
#include <stdlib.h>
#include <string.h>
*/
import "C"

import (
	"unsafe"

	"github.com/unmango/gnumake-go"
)

func Test(nm string, argc uint32, argv [][]byte) *byte {
	mem := gnumake.Alloc(4)
	// C.strncpy((*C.char)(buf), C.CString(""), 4)
	// fmt.Fprintln(os.Stderr, "Test log")
	buf := C.GoBytes(unsafe.Pointer(mem), 4)
	copy(buf, "test")
	return mem
}

func RegisterFuncs() {
	gnumake.AddFunction("testfunc", Test, 0, 0, 0)
}
