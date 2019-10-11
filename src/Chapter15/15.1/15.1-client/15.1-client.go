package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"os"
	"strings"
)

func StartClient(){
	// open connection
	conn,err:=net.Dial("tcp","localhost:50000")
	if err!=nil{
		log.Fatalf("Error dialing, %v",err.Error())
		return
	}
	inputReader:=bufio.NewReader(os.Stdin)
	fmt.Println("what is your name?")
	clientName,_:=inputReader.ReadString('\n')
	log.Printf("clientName is %v",clientName)
	trimmedClient:=strings.Trim(clientName,"\r\n")
	//send info to server until Quit
	for{
		fmt.Println("want to send to the server? Type Q to quit.")
		input,_:=inputReader.ReadString('\n')
		trimmedInput:=strings.Trim(input,"\r\n")
		if trimmedInput=="Q"{
			return
		}
		_,err:=conn.Write([]byte(trimmedClient+" says: "+trimmedInput))
		if err!=nil {
			log.Fatalf("write error: %v",err.Error())
		}
	}
}

func main(){
	StartClient()
}
