package parser

import (
	"learngo/spider/engine"
	"regexp"
)

const cityRe = `<a href="http://album.zhenai.com/u/([0-9]+)"[^>]*>([^<]+)</a>`; //个人用户profile url

func ParseCity(contents []byte) engine.ParseResult  {
	re := regexp.MustCompile(cityRe)
	matchs := re.FindAllSubmatch(contents,-1)
	result := engine.ParseResult{}

	for _, m := range matchs{
		name := string(m[2])
		result.Items = append(result.Items,string(m[2]))
		result.Reuqests = append(result.Reuqests,engine.Request{
			Url: "http://m.zhenai.com/u/" + string(m[1]), //pc版的跑不通
			ParserFunc: func(c []byte) engine.ParseResult {
				return ParseProfile(c,name)
			},
		})
	}
	return result
}