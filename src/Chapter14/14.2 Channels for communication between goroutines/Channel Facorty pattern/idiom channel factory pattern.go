package main

import (
	"fmt"
	"time"
)

// the factory produce products , the products are sent to channel
// the consumer get their products from the channel
// this is just life consensus

// make the channel and return it like a factory

func main() {
	go suck(pump())
	time.Sleep(time.Second)
}
func pump() chan int {
	ch := make(chan int)
	go func() {
		for i := 0; i < 100; i++ {
			ch <- i
		}
	}()
	return ch
	// the channel will be passed to suck , the pump is executed just one time, but the goroutine won't stop until i = 99
}
func suck(ch chan int) {
	for {
		fmt.Println(<-ch)
	}
}
