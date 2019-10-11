package Chapter14

import "log"

func E9Random_bitgen(){
	bit1:=make(chan byte)
	bit0:=make(chan byte)
	go genBit0(bit0)
	go genBit1(bit1)
	var Number []byte
	for i:=0;i<100;i++{
		select {
		case b1:=<-bit1:
			Number=append(Number,b1)
		case b0:=<-bit0:
			Number=append(Number,b0)
		}
	}
	log.Printf("BitGen: %v",Number)
}
func genBit1(bit1 chan byte){
	for{
		bit1<-1
	}
}
func genBit0(bit0 chan byte){
	for{
		bit0<-0
	}
}
