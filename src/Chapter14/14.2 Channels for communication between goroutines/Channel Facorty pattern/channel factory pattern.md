# channel facotry pattern
工厂生产往channel中填入数据，客户从channel中读取数据
````go
func main(){
	go suck(pump())
	time.Sleep(time.Second)
}
````
## 生产数据的工厂
````go
func pump()chan int{
	ch:=make(chan int)
	go func(){
		for i:=0;i<100 ; i++{
			ch<-i
		}
	}()
	return ch
}
````
`pump()`执行一次，该线程就停止了，但其中的goroutine不会停下，由于匿名函数捕捉了`ch`，因此`ch`也不会被回收掉。这样实际上产生数据的只是那个goroutine，而`pump()`只是起到了`ch:=make(chan int)`与`return ch`的作用后就结束其生命周期
## 不断拉取数据的客户
````go
func suck(ch chan int){
	for {
		fmt.Println(<-ch)
	}
}
````
该协程会一直存在，直到`mian()`退出，而且当`<-ch`无数据时，受到阻塞

