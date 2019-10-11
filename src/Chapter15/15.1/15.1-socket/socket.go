package main

import (
	"fmt"
	"io"
	"net"
)

func main() {
	var (
		host          = "www.apache.org"
		port          = "80"
		remote        = host + ":" + port
		msg    string = "GET / \n"
		data          = make([]uint8, 4096)
		read          = true
		count         = 0
	)
	// create the socket
	con, err := net.Dial("tcp", remote)
	// send message an HTTP GET request
	_, _ = io.WriteString(con, msg)
	// read the response from the webserver
	for read {
		count, err = con.Read(data)
		read = err == nil
		fmt.Printf(string(data[0:count]))
	}
	_ = con.Close()
}
