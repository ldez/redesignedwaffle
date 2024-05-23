package main

import "golang.org/x/exp/maps"

func Boo() {
	d := make(map[string]int)
	maps.Keys(d)
}
