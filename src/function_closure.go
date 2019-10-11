package main

import (
	"fmt"
	"log"
	"strings"
)

func test_keep_local_var_in_closureFun() {
	var f = Adder2()
	f(1)
	f(20)
	f(300)
}
func Adder2() func(int) int {
	var x int
	return func(delta int) int {
		x += delta
		defer func() {
			fmt.Printf("f(%d)=%d\n", delta, x)
		}()
		return x
	}
}
func MakeAddSuffix(suffix string) func(string) string {
	return func(name string) (res string) {
		orianaName := name
		res = name
		defer func() {
			log.Printf("Add Orignal Name %s with Suffix %s and get %s\n", orianaName, suffix, res)
		}()
		if !strings.HasSuffix(name, suffix) {
			res += suffix
			return res
		}
		return res
	}
}
func testAddSuffix_closure_fun() {
	addBmp := MakeAddSuffix(".bmp")
	addJPG := MakeAddSuffix(".jpg")
	addBmp("file1")
	addJPG("file2")

}
