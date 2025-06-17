package main

import "github.com/unmango/gnumake-go"

//export InitMkFunc
func InitMkFunc() {
	gnumake.AddFunction("", func(nm string, argc uint32, argv [][]byte) *byte {
		return nil
	}, 0, 0, 0)
}

func main() {}
