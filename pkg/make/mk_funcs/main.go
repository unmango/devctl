package main

import "C"

import (
	"fmt"

	"github.com/unmango/gnumake-go"
)

//export InitMkFunc
func InitMkFunc() {
	gnumake.AddFunction("test", func(nm string, argc uint32, argv [][]byte) *byte {
		fmt.Println("called custom function")
		return nil
	}, 0, 0, 0)
}

func main() {}
