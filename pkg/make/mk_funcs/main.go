package main

import "github.com/unmango/gnumake-go"

/*
int plugin_is_GPL_compatible;
*/

func InitMkFunc() {
	gnumake.AddFunction("", func(nm string, argc uint32, argv [][]byte) *byte {
		return nil
	}, 0, 0, 0)
}

func main() {}
