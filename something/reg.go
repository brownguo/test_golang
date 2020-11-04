package main

import (
	"fmt"
	"regexp"
)

const text = `aa@qq.com
	aa@aa.com
`
const cityRe = `<a href="(http://album.zhenai.com/u/[0-9]+)"[^>]*>([^<]+)</a>`;
const profileRe = `<span class="nickName"[^>]*>([^<]+)</span>`

func reProfile ()  {
	re := regexp.MustCompile(profileRe)
	d := `<span class="nickName" data-v-3c42fade>燕子</span>`
	submatch := re.FindAllStringSubmatch(d, -1)

	for _,m := range submatch{
		fmt.Println(m[1])
	}
}

func reCity()  {
	re := regexp.MustCompile(cityRe)
	d := `<a href="http://album.zhenai.com/u/1012253053" target="_blank">漂流瓶</a>`
	submatch := re.FindAllStringSubmatch(d,-1)

	fmt.Println(submatch)
}

func main() {
	reProfile()
	return
	re := regexp.MustCompile(`([a-zA-Z0-9]+)@([a-zA-Z0-9]+)(\.[a-zA-Z0-9.]+)`)
	match := re.FindAllStringSubmatch(text,-1)
	for _,m := range match{
		fmt.Println(m)
	}
}
