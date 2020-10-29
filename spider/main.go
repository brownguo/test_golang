package main

import (
	"learngo/spider/engine"
	"learngo/spider/zhenai/parser"
)

func main() {
	url := "https://www.zhenai.com/zhenghun"  //城市列表页面
	engine.SimpleEngine{}.Run(engine.Request{
		Url: url,
		ParserFunc: parser.ParseCityList,
	})
}