package main

import "fmt"

const (
	AvailableMemory         = 10 << 20 // 10 MB
	AverageMemoryPerRequest = 10 << 10 // 10KB
	MaxREQS                 = AvailableMemory / AverageMemoryPerRequest
)

var sem = make(chan int, MaxREQS)

type Request struct {
	a, b   int
	replyc chan int
}

func process(r *Request) {
	r.replyc <- r.a + r.b
}

func handler(r *Request) {
	process(r)
	fmt.Printf("handle done at req: %+v\n", r)
	<-sem // done make 1 empty place in the buffer
}

func Sever(queue chan *Request) {
	for {
		sem <- 1
		//sem <- 1 // take a place for handler
		// blocks when channel is full
		// wait until there is capacity to process a request
		request := <-queue
		go handler(request) // start each goroutine for each request which implements replyc
	}
}

func main() {
	queue := make(chan *Request)
	go Sever(queue)

	const requestNum = MaxREQS + 1
	var reqs [requestNum]Request
	for i := 0; i < requestNum; i++ {
		req := &reqs[i]
		req.a = i
		req.b = i
		req.replyc = make(chan int)
		queue <- req
		fmt.Printf("handle request %d-th\n", i)
	}

	// checks:
	for i := requestNum - 1; i >= 0; i-- {
		if <-reqs[i].replyc != i*2 {
			fmt.Println("fail at", i)
		} else {
			fmt.Println("Request ", i, " is ok!")
		}
	}

}
