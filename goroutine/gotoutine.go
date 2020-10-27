package main

import (
	"fmt"
	"time"
)

func main() {
	//var s [10]int

	for i := 0; i < 10; i++ {
		go func(flag int) {
			for  {
				//s[flag] ++
				fmt.Printf("Hello form goroutine:%d \n",flag)
			}
		}(i)
	}
	time.Sleep(time.Microsecond)
	//fmt.Println(s)
}
