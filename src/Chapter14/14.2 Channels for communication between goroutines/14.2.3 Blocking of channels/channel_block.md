# channel block
````go
func pump(ch chan int){
	for i:=0; ;i++{
		ch<-i
	}
}

func main() {
	ch1:=make(chan int)
	go pump(ch1) // pump hang
	fmt.Println(<-ch1)
	fmt.Println(<-ch1)
}

```` 
将输出 0 1
#导致死锁的情况
**死锁就是主线程遭到了阻塞，协程阻塞与否不构成死锁**
阻塞的产生有两种情况：
1. 写入channel时，channel没有多余的空间，会阻塞至接收者拿取之；因此主线程在写入时，要确保读取方先启动了，否则会阻塞主线程，导致接收方无法读取；
2. 读取时，channel没有数据，会阻塞至发送者发送消息；因此主线程在读取channel时，要确保发送方协程先启动了，否则阻塞主线程，导致死锁。
## 1. 没有启动接收方，死锁
````go
func main() {
	ch1:=make(chan int)
	fmt.Println(<-ch1)
}
````
## 2. 接收方未先启动就写入，死锁
如何通道已经充满了数据，继续向通道发送数据`ch1<-data`，则会导致发送方阻塞
````go
func main() {
	ch1:=make(chan int)
	ch1<-1
	fmt.Println(<-ch1)
}
````
第三行阻塞主线程，导致死锁
```` go
func main() {
	ch1:=make(chan int)
	ch1<-2
	go func() {
		fmt.Println(<-ch1)
	}()
}
````
同理，第三行阻塞，协程未创建
## 解决方法
先启动读channel协程，在写入channel
````go
func main() {
	ch1:=make(chan int)
	go func(){
		fmt.Println(<-ch1)
	}()
	ch1<-2
}
````
先启动写channel协程，在读channel
````go
func main() {
	ch1:=make(chan int)
	go func(){
		ch1<-2
	}()
	fmt.Println(<-ch1)
}
````
# 总结
阻塞主线程就时死锁，其他协程被阻塞，可以通过主线程的调度来恢复，但主线程一旦阻塞就没救了。