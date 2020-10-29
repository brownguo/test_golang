package parser

import (
	"learngo/spider/engine"
	"learngo/spider/model"
	"regexp"
	"strconv"
)

var descRe  = regexp.MustCompile(`<p data-v-77281b7a>([^<]+)</p>`)
var ageRe	= regexp.MustCompile(`<div class="tag" data-v-10352ec0>([\d]+)岁</div>`)

func ParseProfile(contents []byte,name string) engine.ParseResult  {
	profile := model.Profile{}
	//昵称
	profile.NickName = name
	//个性签名
	profile.Desc = extractString(contents,descRe)
	//年龄
	age,err := strconv.Atoi(extractString(contents,ageRe))
	if err == nil{
		profile.Age = age
	}
	result := engine.ParseResult{
		Items: []interface{}{profile},
	}
	return result
}

func extractString(contents []byte,re *regexp.Regexp) string {
	match := re.FindSubmatch(contents)
	if len(match) >= 2{
		return string(match[1])
	}else{
		return ""
	}
}