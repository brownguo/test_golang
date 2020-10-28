package fetcher

import (
	"bufio"
	"fmt"
	"golang.org/x/net/html/charset"
	"golang.org/x/text/encoding"
	"golang.org/x/text/encoding/unicode"
	"golang.org/x/text/transform"
	"io"
	"io/ioutil"
	"log"
	"net/http"
)

func Fetch(url string) ([]byte,error)  {
	client := &http.Client{}
	request, err := http.NewRequest("GET", url, nil)
	request.Header.Set("user-agent","Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_5) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/86.0.4240.111 Safari/537.36")
	if err != nil{
		panic(err)
	}

	response, _ := client.Do(request)

	defer response.Body.Close()

	if response.StatusCode != http.StatusOK{
		return nil,fmt.Errorf(" 1wrong status code:%d",response.StatusCode)
	}
	//转码
	e := determineEncoding(response.Body)
	utf8Reader := transform.NewReader(response.Body,e.NewDecoder())
	return ioutil.ReadAll(utf8Reader)
}

func determineEncoding(r io.Reader) encoding.Encoding {
	bytes, err := bufio.NewReader(r).Peek(1024)
	if err != nil{
		log.Printf("fetch error %v:",err)
		return unicode.UTF8
	}
	e, _, _ := charset.DetermineEncoding(bytes, "")
	return e
}
