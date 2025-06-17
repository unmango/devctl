package main

/*
int plugin_is_GPL_compatible;
*/
import "C"

import "github.com/unmango/gnumake-go"

func InitMkFunc() {
	gnumake.AddFunction("", func(nm string, argc uint32, argv [][]byte) *byte {
		return nil
	}, 0, 0, 0)
}

func main() {}
