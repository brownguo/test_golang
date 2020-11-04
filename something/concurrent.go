package main

import "fmt"

func main() {
	//创建存放int类型的通道
	c := make(chan int)
	//子协程
	go func() {
		defer fmt.Println("子协程结束")
		fmt.Println("子协程正在运行...")
		//将666发送到通道c
		c <- 666
	}()
	//若已取出数据，下面再取会报错
	a := <-c
	fmt.Println("a = ", a)
	//主协程取数据
	//从c中取数据
	num := <-c
	fmt.Println("num = ", num)
	fmt.Println("主协程结束")
}