package proxy

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
)

//GetProxy 获取一个私密代理
func GetProxy(*http.Request) (*url.URL, error) {
	URL := "https://dps.kdlapi.com/api/getdps/?orderid=983269068917124&num=1"
	resp, err := http.Get(URL)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	text, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	username := "2227627947"
	password := "nl44w92g"
	proxyURL := fmt.Sprintf("https://%s:%s@%s", username, password, text)
	fmt.Println(proxyURL)
	return url.Parse(proxyURL)
}

//GetProxy2 获取独享代理
func GetProxy2(*http.Request) (*url.URL, error) {
	return url.Parse("http://117.34.124.81:16816")
}
