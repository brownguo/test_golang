package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
)


func main() {
	var id int
	var address string

	db, err := sql.Open("mysql", "root:Admin123.@/WB_USER?charset=utf8")
	if err != nil{
		panic(err)
	}

	rows, err := db.Query("select id,address from address limit 10")

	defer db.Close()
	if err != nil{
		panic(err)
	}
	for rows.Next(){ //迭代查询数据
		rows.Scan(&id,&address) //读取每一行数据
		fmt.Println(id,address)
	}

}
