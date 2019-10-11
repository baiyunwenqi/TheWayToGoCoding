package main

import (
	"log"
	"net"
)

type server struct {
	ClientNames []string
	CloseFlag   bool
}

func main() {
	var host = "locahost:8000"
	log.Printf("Starting the sever...")
	listener, err := net.Listen("tcp", host)
	checkError(err, "listing")
	// listen and accept connections from clients:
	Notclose := true
	for Notclose {
		conn, err := listener.Accept()
		checkError(err, "accept")
		go doServerStuff(conn)
		if !Notclose {
			log.Printf("close the server")
		}
	}
}
func NewServer() *server {
	ser := &server{
		ClientNames: make([]string, 0),
		CloseFlag:   false,
	}
	return ser
}
func doServerStuff(conn net.Conn) {
	for {
		buf := make([]byte, 512)
		_, err := conn.Read(buf)
		if err != nil {
			log.Printf("Error reading,%v", err.Error())
		}
		log.Printf("Received data: %v", string(buf))

	}
}
func checkError(err error, info string) {
	if err != nil {
		panic("ERROR: " + info + " " + err.Error())
	}
}
