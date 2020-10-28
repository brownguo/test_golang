package parser

import (
	"learngo/spider/engine"
	"learngo/spider/model"
	"regexp"
)

var descRe  = regexp.MustCompile(`<p data-v-77281b7a>([^<]+)</p>`)
func ParseProfile(contents []byte) engine.ParseResult  {
	profile := model.Profile{}

	descMatchs := descRe.FindSubmatch(contents)
	if descMatchs != nil{
		profile.Desc = string(descMatchs[1])
	}

	result := engine.ParseResult{
		Items: []interface{}{profile},
	}
	return result
}