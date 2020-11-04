package main

import (
	"fmt"
	"time"
)

type task struct {
	taskList []byte
}

func main() {
	 task := make(chan int)
	 fmt.Printf("len(c)=%d, cap(c)=%d\n", len(task), cap(task))
	//子协程存数据
	go func() {
		defer fmt.Println("子协程结束")
		//向通道添加数据
		for i := 0; i < 3; i++ {
			task <- i
			fmt.Printf("子协程正在运行[%d]:len(c)=%d,cap(c)=%d\n", i, len(task), cap(task))
		}
	}()
	time.Sleep(2 * time.Second)
	//主协程取数据
	for i := 0; i < 3; i++ {
		task := <-task
		fmt.Println("task=", task)
	}
	fmt.Println("主协程结束")
}