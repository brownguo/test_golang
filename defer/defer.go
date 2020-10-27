package main

import (
	"bufio"
	"fmt"
	"learngo/defer/fib"
	"os"
)

func tryDefer() {
	defer fmt.Println("after ....")
	fmt.Println("befor print!")
}


func writeFibToFile(filename string)  {
	file,err := os.Create(filename)
	if err != nil{
		panic(err)
	}
	defer file.Close()
	writer := bufio.NewWriter(file) //搞一个buffer ，先写到内存里，然后才用内存移动到文件。
	defer writer.Flush()
	f := fib.Fibcci()
	for i:=0;i<20;i++ {
		fmt.Fprintln(writer,f())
	}
}

func main() {
	tryDefer()
	msg := fib.PrintMsg()
	fmt.Println(msg())
	writeFibToFile("dadada.txt")
}
