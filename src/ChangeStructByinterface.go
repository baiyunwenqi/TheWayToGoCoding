package main

import "fmt"

type Cp interface {
	ChangeValue(int, int)
}

type complexType struct {
	real int
	imag int
}

func (c *complexType) ChangeValue(newR int, newI int) {
	c.real = newR
	c.imag = newI
}
func Change(cp Cp) {
	cp.ChangeValue(1, 8)
}
func testChange() {
	c := complexType{1, 3}
	Change(&c)
	fmt.Printf("%v", c)
}
