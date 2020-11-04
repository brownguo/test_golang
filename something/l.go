package main

import (
	"fmt"
	"net/http"
	"net/http/httputil"
)
func main() {
	var url = "http://sir.hitik.net:8080/api.php?t=1603865099805"
	request,err := http.NewRequest(http.MethodPost,url,nil)

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