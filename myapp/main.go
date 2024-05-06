package main

import (
	"errors"
	"fmt"
)

func main() {
	fmt.Println("Hello")
	foo("a")
	fmt.Println("world")
	bar("a")
}

func foo(a string) error {
	if a == "a" {
		return nil
	}

	return errors.New("OOPS")
}

func bar(b string) error {
	if b == "b" {
		return nil
	}

	return errors.New("OOPS I did it again")
}
