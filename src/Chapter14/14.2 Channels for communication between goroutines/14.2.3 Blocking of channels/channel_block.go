package main

import (
	"fmt"
)

func pump(ch chan int) {
	for i := 0; ; i++ {
		ch <- i
	}
}

func main() {
	ch1 := make(chan int)
	go func() {
		ch1 <- 2
	}()
	fmt.Println(<-ch1)
}
