package main

import "fmt"

func fibcci() func() int {
	a,b := 0,1
	return func() int {
		a, b = b, a+b
		return a
	}
}

func main() {
	f := fibcci()
	for i:=0; i<10; i++{
		fmt.Println(f())
	}
}
