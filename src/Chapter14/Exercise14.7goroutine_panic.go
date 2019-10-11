package Chapter14

import (
	"fmt"
	"log"
)

func tel(ch chan int){
	for i:=0;i<15;i++{
		ch<-i
	}
	close(ch)
}
func E7Goroutine_Panic(){
	ch:=make(chan int)
	go tel(ch)
	for{
		if in,ok:=<-ch;ok {
			fmt.Printf("Get Number: %d\n",in)
		}
		break
	}
}

func E7Goroutine_close(){
	ch:=make(chan int)
	go tel(ch)
	for in:=range ch{
			fmt.Printf("Get Number: %d\n",in)
	}
}

func telNoticEnd(ch chan int,closed chan bool){
	for i:=0;i<15;i++{
		ch<-i
	}
	closed<-true
	//close(ch)
}
func E7Goroutine_select(){
	ch:=make(chan int)
	closed:=make(chan bool)
	go telNoticEnd(ch,closed)
	CanRun:=true
	for CanRun{
		select {
		case number,ok:=<-ch:
			log.Printf("check ch1")
			if ok{
				log.Printf("Get Number: %d\n",number)
			}
		case clos:=<-closed:
			log.Printf("Close : %v\n",clos)
			if clos{
				//os.Exit(0)
				CanRun=false
			}
		}
	}
}