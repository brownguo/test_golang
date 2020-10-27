package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httputil"
)

func httpGet()  {
	resp,err := http.Get("http://api.moeya.cn/")
	if err != nil{
		panic(err)
	}
	defer resp.Body.Close()
	body,err := ioutil.ReadAll(resp.Body)
	fmt.Println(string(body))
}

func main() {
	httpGet()
	var url = "https://api.map.baidu.com/?qt=gc&wd=%E5%8C%97%E4%BA%AC%E5%B8%82%E7%9F%B3%E6%99%AF%E5%B1%B1%E5%8C%BA%E5%85%AB%E8%A7%92%E5%8D%97%E9%87%8C%E7%94%B21%E5%8F%B7%E6%A5%BC&cn=%E5%8C%97%E4%BA%AC%E5%B8%82&ie=utf-8&oue=1&fromproduct=jsapi&res=api&ak=CG8eakl6UTlEb1OakeWYvofh"
	request,err := http.NewRequest(http.MethodGet,url,nil)

	request.Header.Add("User-Agent",
		"Mozilla/5.0 (iPhone; CPU iPhone OS 10_3 like Mac OS X) AppleWebKit/602.1.50 (KHTML, like Gecko) CriOS/56.0.2924.75 Mobile/14E5239e Safari/602.1")
	clent := http.Client{}
	resp,err := clent.Do(request)
	if err != nil{
		panic(err)
	}
	defer resp.Body.Close()
	s,err := httputil.DumpResponse(resp,true)
	if err != nil{
		panic(err)
	}
	fmt.Printf("%s\n",s)
}
