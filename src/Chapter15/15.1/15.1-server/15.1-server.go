package main

import (
	"log"
	"net"
)

func main(){
	log.Printf("Starting the sever...")
	// create listener:
	listener,err:=net.Listen("tcp","localhost:50000")
	if err!=nil{
		log.Fatalf("Error listening %v",err.Error())
		return
	}
	// listen and accept connections from clients:
	for{
		conn,err:=listener.Accept();
		if err!=nil {
			log.Fatalf("Error accepting %v", err.Error())
			return
		}
			go doServerStuff(conn)
		}
}
func doServerStuff(conn net.Conn){
	for{
		buf:=make([]byte,512)
		_,err:=conn.Read(buf)
		if err!=nil{
			log.Printf("Error reading,%v",err.Error())
		}
		log.Printf("Received data: %v",string(buf))
	}
}

