package main

import (
	"C"

	"github.com/unmango/devctl/pkg/make"
)

//export InitMkFunc
func InitMkFunc() {
	make.RegisterFuncs()
}

func main() {}
