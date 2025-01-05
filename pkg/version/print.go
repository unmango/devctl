package version

import "fmt"

func Cat(path string) (int, error) {
	if v, err := ReadFile(path); err != nil {
		return 0, err
	} else {
		return fmt.Print(v)
	}
}

func Print(name string) (int, error) {
	return Cat(RelPath(name))
}
