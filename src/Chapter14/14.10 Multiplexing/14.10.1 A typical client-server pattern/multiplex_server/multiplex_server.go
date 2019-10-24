package main

import "fmt"

type Request struct {
	a, b   int
	replyc chan int // reply channel inside the Request
}

type binOp func(a, b int) int

func run(op binOp, req *Request) {
	req.replyc <- op(req.a, req.b) // send back through replyc channel
}

func server(op binOp, service chan *Request) {
	for {
		req := <-service // requests arrive here
		go run(op, req)  // each req start a different goroutine go means new goroutine
	}
}
func startServer(op binOp) chan *Request {
	reqChan := make(chan *Request)
	go server(op, reqChan) // server on reqChan
	return reqChan         // communicate through reChan with client
}

func main() {
	adder := startServer(func(a, b int) int { return a + b })
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
	fmt.Println("done")
}
