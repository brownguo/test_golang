package main

import (
	"fmt"
	"time"
)

func main() {
	var s [3]int

	for i := 0; i < 3; i++ {
		go func(flag int) {
			for  {
				s[flag] ++
				fmt.Printf("Hello form goroutine:%d \n",flag)
			}
		}(i)
	}
	time.Sleep(100 * time.Microsecond)
}
