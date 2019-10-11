package main

import (
	"fmt"
	"log"
	"runtime"
)

func test_Closure_fun_debug() {
	where := func() {
		_, file, line, _ := runtime.Caller(1)
		log.Printf("%s:%d", file, line)
	}
	where()
	fmt.Print("sdfs\n")
}
