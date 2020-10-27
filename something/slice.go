package main

import "fmt"

func printSlice(s []int)  {
	fmt.Printf("%v,len:%d,cap:%d \n",s,len(s),cap(s))
}

func appendSlice()  {
	var s = []int{}
	for i:=0 ; i< 100 ; i++ {
		s = append(s, i*2+1)
	}
	fmt.Println(s)
}

func main() {
	arr1 := [...]int {0,1,2,3,4,5,6,7,8,9}
	//切片扩展的时候数组不能超过cap 
	fmt.Printf("len:%d,cap:%d \n", len(arr1),cap(arr1))
	//make一个切片，带cap长度的。
	s := make([]int,10,12)
	printSlice(s)
	appendSlice()
}
