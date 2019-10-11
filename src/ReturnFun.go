package main

import "fmt"

func Add2() func(b int) int {
	return func(b int) int {
		return b + 2
	}
}
func Adder(a int) func(b int) int {
	return func(b int) int {
		return a + b
	}
}
func testReturnFun() {
	p2 := Add2()
	fmt.Printf("Call Add2 for 3 gives: %v\n", p2(3))
	TwoAdder := Adder(2)
	fmt.Printf("The result using Adder is :%v\n", TwoAdder((3)))
}
