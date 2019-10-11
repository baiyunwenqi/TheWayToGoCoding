package main

import "fmt"

func textExpendSlice() {
	oldSlice := make([]int, 3, 10)
	newSlice := ExpandSlice(4, oldSlice)
	fmt.Printf("old slice len is %d\nnew slice len is %d\n", len(oldSlice), len(newSlice))
	fmt.Printf("old slice cap is %d\nnew slice cap is %d\n", cap(oldSlice), cap(newSlice))
}
func ExpandSlice(factor int, slice []int) []int {
	lent := len(slice)
	newL := lent * factor
	if newL > cap(slice) {
		newSlice := make([]int, newL, newL*2)
		copy(newSlice, slice)
		slice = newSlice
	}
	slice = slice[0:newL]
	return slice
}
