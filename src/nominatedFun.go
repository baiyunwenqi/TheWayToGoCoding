package main

import "fmt"

func testnominatedFun() {
	fv := func(s string) { fmt.Printf("%s\n", s) }
	fv("hello word")
}
