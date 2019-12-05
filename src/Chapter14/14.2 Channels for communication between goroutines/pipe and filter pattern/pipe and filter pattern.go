package main

import (
	"fmt"
	"log"
)

func generate(ch chan int){ // generate prime
	for i:=2;i<12;i++{
		ch<-i
		//time.Sleep(time.Second)
	}
	close(ch)
	log.Println("close ch")
}

func filter(in,out chan int, prime int){
	for {
		i,ok:=<-in // receive value from main()
		if !ok{
			log.Println("ch has been closed")
			close(out)
			return
		}
		if i%prime!=0{
			out<-i
		}
	}
}

func main(){
	ch:=make(chan int)
	go generate(ch)
	for{
		prime,ok:=<-ch
		if !ok{
			return
		}
		fmt.Print(prime," ")
		ch1:=make(chan int)
		go filter(ch,ch1,prime)
		i,isClose:=<-ch1
		if !isClose{
			break
		}else{
			go func() {
				ch1<-i
			}()
			ch=ch1
		}
	}
}