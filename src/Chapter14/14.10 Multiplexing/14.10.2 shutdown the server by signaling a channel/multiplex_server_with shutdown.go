package main

import (
	"fmt"
	"github.com/pantsing/log"
)

func init() {
	log.SetFlags(log.Lmodule | log.Llevel | log.Ldate | log.Lmicroseconds | log.Lshortfile)
	log.Std.Level = log.Ldebug
}

type Request struct {
	a, b   int
	replyc chan int // reply channel inside the Request
}

type binOp func(a, b int) int

func run(op binOp, req *Request) {
	req.replyc <- op(req.a, req.b) // send back through replyc channel
}

func server(op binOp, service chan *Request, quit chan bool) {
	for {
		select {
		case req := <-service: // requests arrive here
			go run(op, req) // each req start a different goroutine go means new goroutine
		case <-quit:
			log.Info("receive quit signal")
			fmt.Println("receive quit signal")
			return
		}
	}
}
func startServer(op binOp) (chan *Request, chan bool) {
	reqChan := make(chan *Request)
	quit := make(chan bool)
	go server(op, reqChan, quit) // server on reqChan
	return reqChan, quit         // communicate through reChan with client
}

func main() {
	adder, quit := startServer(func(a, b int) int { return a + b })
	const N = 100
	var reqs [N]Request
	for i := 0; i < N; i++ {
		req := &reqs[i]
		req.a = i
		req.b = i + N
		req.replyc = make(chan int)
		adder <- req // once req is passed to run(op,req), a new goroutine is established to handle the request, and sent back through req too
	}
	// checks:
	for i := N - 1; i >= 0; i-- {
		if <-reqs[i].replyc != N+2*i {
			fmt.Println("fail at", i)
		} else {
			fmt.Println("Request ", i, " is ok!")
		}
	}
	quit <- true
	fmt.Println("done")
}
