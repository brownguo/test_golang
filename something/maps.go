package main

import "fmt"

func main() {
	m := map[string]string{
		"name":"yexuanguo",
		"age":"18",
		"gender":"1",
	}
	info := map[string]string{
		"1":"男",
		"2":"女",
	}
	fmt.Printf("name is : %s \n",m["name"])
	fmt.Printf("age is : %s \n",m["age"])
	fmt.Printf("gender is : %s \n",info[m["gender"]])
	fmt.Println(m)

	//声明两个map
	var m1 = make(map[string] int)
	m2 := make(map[int]int)
	fmt.Println(m1)
	fmt.Println(m2)
}
