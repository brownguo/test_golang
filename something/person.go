package main

import "fmt"

type Person struct {
	Name string
	Age int
	Desc string
}

func (info *Person) GetPerson()  {
	fmt.Printf("Name:%s,Age:%d",info.Name,info.Age)
}
func (info *Person) GetDesc()  {
	info.Desc = "因为Persoon传的是指针，改了一下DESC！"
	fmt.Printf("Desc:%s",info.Desc)
}