package main

import (
	"errors"
	"fmt"

	"golang.org/x/exp/maps"
	"gopkg.in/yaml.v3"
)

func main() {
	fmt.Println("Hello")
	bar("a")
	fmt.Println("world")
	foo("b")
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

func Boo() {
	yaml.Marshal(nil)
	d := make(map[string]int)
	maps.Keys(d)
}
