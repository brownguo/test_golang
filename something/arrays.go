package main

import "fmt"

func main() {
	arr1 := [...]int {2,4,6,8,10}
	for i := 0 ; i<len(arr1); i++ {
		fmt.Println(arr1[i])
	}
	for i,v := range arr1{
		fmt.Printf("index:%s,value:%s \n", i,v)
		fmt.Println(i,"===>",v)
	}
}
