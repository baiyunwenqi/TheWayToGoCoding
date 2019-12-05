package main

import (
	"fmt"
	"sort"
)

// in this snippet we start 2 goroutines , each sort a part of a slice, anc send signal to shut down the main process

func main() {
	done := make(chan bool)
	s := []int{5, 8, 9, 2, 3, 4, 8, 7}
	doSort := func(s []int) {
		sort.Ints(s)
		done <- true
	}
	i := len(s) >> 1
	go doSort(s[:i])
	go doSort(s[i:])
	<-done // release the channel
	<-done
	fmt.Printf("sorted data: %v", s)
}
