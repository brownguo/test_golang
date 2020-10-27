package main

import (
	"fmt"
	"time"
)

func worker(id int, c chan int)  {
	for n := range c{
		fmt.Printf("WorkerID: %d Recv:%c\n",id,n)
	}
}

func createWorker(id int) chan<- int {  //返回chan int  外面只能发数据
	c := make(chan int)
	go worker(id,c)
	return c
}
func chanDemo()  {
	var channels [10]chan<- int // 只能写 send only.

	for i := 0; i < 10; i++ {
		channels[i] = createWorker(i)  //创建一个worker，返回chan int
	}
	for i := 0; i < 10; i++ {			//往返回的worker里写数据，也就是往channel里写数据
		channels[i] <- 'A' + i
	}
	//for i := 0; i < 10; i++ {			//往返回的worker里写数据，也就是往channel里写数据
	//	channels[i] <- 'a' + i
	//}
	time.Sleep(time.Microsecond)
}

func main() {
	chanDemo()
}

