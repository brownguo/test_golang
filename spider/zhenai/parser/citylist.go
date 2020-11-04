package parser

import (
	"learngo/spider/engine"
	"regexp"
)

const cityListRe = `<a href="(http://www.zhenai.com/zhenghun/[0-9a-z]+)"[^>]*>([^<]+)</a>`  //城市列表

func ParseCityList(contents []byte) engine.ParseResult{

	re := regexp.MustCompile(cityListRe)
	matchs := re.FindAllSubmatch(contents,-1)
	result := engine.ParseResult{}


	limit := 3

	for _, m := range matchs{
		result.Items = append(result.Items,string(m[2])) //城市名字
		result.Reuqests = append(result.Reuqests,engine.Request{
			Url:string(m[1]),
			ParserFunc: ParseCity,
		})
		limit--
		if limit == 0{
			break
		}
	}

	return result
}