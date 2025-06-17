package make

import "github.com/unmango/gnumake-go"

func Test(nm string, argc uint32, argv [][]byte) *byte {
	return nil
}

func RegisterFuncs() {
	gnumake.AddFunction("test", Test, 0, 0, 0)
}
