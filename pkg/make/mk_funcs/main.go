package main

import "github.com/unmango/gnumake-go"

//export plugin_is_GPL_compatible
var plugin_is_GPL_compatible int

func InitMkFunc() {
	gnumake.AddFunction("", func(nm string, argc uint32, argv [][]byte) *byte {
		return nil
	}, 0, 0, 0)
}

func main() {}
