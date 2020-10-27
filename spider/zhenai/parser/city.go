package parser

import (
	"learngo/spider/engine"
	"regexp"
)

const cityRe = `<a href="(http://album.zhenai.com/u/[0-9]+)"[^>]*>([^<]+)</a>`;

func ParseCity(contents []byte) engine.ParseResult  {
	re := regexp.MustCompile(cityRe)
	matchs := re.FindAllSubmatch(contents,-1)
	result := engine.ParseResult{}

	for _, m := range matchs{
		result.Items = append(result.Items,string(m[2]))
		result.Reuqests = append(result.Reuqests,engine.Request{
			Url:string(m[1]),
			ParserFunc: engine.NilParser,
		})
	}
	return result
}