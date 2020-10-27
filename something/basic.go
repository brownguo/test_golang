package main

import "fmt"

var (
	e = "eee"
	f = "fff"
)

const (
	filename = "dadada.js"
)

func varibleInit()  {
	a,b,c,d := 1,2,3,true
	fmt.Println(a,b,c,d)
	h,j := 1,""
	fmt.Printf("%d,%q\n",h, j)
	var name string  = "yexuanguo"
	fmt.Println(name)
}

func main() {
	fmt.Printf("%s, %s ,%s \n\n", e, f,filename)
	varibleInit()
}