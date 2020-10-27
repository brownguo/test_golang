package main

import (
	"fmt"
	"io/ioutil"
)

func calcScore(score int) string  {
	g := ""
	switch  {
	case score > 1:
		g = "> 1"
	case score < 1:
		g = "< 1"
	case score == 1:
		g = "== 1"
	}
	return g
}

func main() {
	const filename = "test.1txt"
	contents, err := ioutil.ReadFile(filename)
	if err != nil{
		fmt.Println(err)
	}else {
		fmt.Printf("%s", contents)
	}
	fmt.Println(calcScore(0123))
}
