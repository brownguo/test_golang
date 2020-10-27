package main

import (
	"fmt"
	"time"
)

func worker(id int, c chan int)  {
	for n := range c{
		fmt.Printf("Worker: %d Recv:%c\n",id,n)
	}
}

func createWorker(id int) chan<- int {  //返回chan int  外面只能发数据
	c := make(chan int)
	go worker(id,c)
	return c
}

func bufferedChannel()  {
	c := make(chan int,3) //3个管道
	go worker(0,c)
	c <- 'a'
	c <- 'b'
	c <- 'c'
	c <- 'd'
	close(c)
	time.Sleep(time.Microsecond)
}

func channelClose()  {
	c := make(chan int)
	go worker(0,c)
	c <- 'a'
	c <- 'b'
	c <- 'c'
	c <- 'd'
	close(c)
	time.Sleep(time.Microsecond)
}

func chanDemo()  {
	var channels [10]chan<- int // 只能写 send only.

	for i := 0; i < 10; i++ {
		channels[i] = createWorker(i)  //创建一个worker，返回chan int
	}
	for i := 0; i < 10; i++ {			//往返回的worker里写数据，也就是往channel里写数据
		channels[i] <- 'A' + i
	}
	time.Sleep(time.Microsecond)
}

func main() {
	fmt.Println("Channel as first-class citizen \n")
	chanDemo()
	fmt.Println("Buffered Channel \n")
	bufferedChannel()
	fmt.Println("Channel close and range \n")
	channelClose()
}

