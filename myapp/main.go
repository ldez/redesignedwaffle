package main

import (
	"errors"
	"fmt"
)

func main() {
	fmt.Println("Hello")
	foo("a")
}

func foo(a string) error {
	if a == "a" {
		return nil
	}

	return errors.New("OOPS")
}
