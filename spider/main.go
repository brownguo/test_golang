package main

import (
	"learngo/spider/engine"
	"learngo/spider/zhenai/parser"
)

func main() {
	url := "https://www.zhenai.com/zhenghun"
	engine.Run(engine.Request{
		Url: url,
		ParserFunc: parser.ParseCityList,
	})
}