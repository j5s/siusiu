package controllers

import (
	"biu/pkg/proxy"
	"biu/validate"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"time"

	"github.com/abiosoft/ishell"
)

//GoogleSearchController  google url 采集
func GoogleSearchController(c *ishell.Context) {
	fmt.Println("google url采集")
	keyword, err := validate.ReadKeyword(c)
	if err != nil {
		log.Println("validate.ReadKeyword failed,err:", err)
		return
	}
	url := fmt.Sprintf("https://www.google.com/search?q=%s", keyword)
	request, err := http.NewRequest("http.MethodGet", url, nil)
	if err != nil {
		log.Println("http.NewRequest faild,err:", err)
		return
	}
	request.Header.Add("User-Agent", proxy.GetUserAgent())
	client := &http.Client{
		Timeout: 3 * time.Second,
		Transport: &http.Transport{
			Proxy: proxy.GetProxy2,
		},
	}
	var resp *http.Response
	for range time.Tick(3 * time.Second) {
		resp, err = client.Do(request)
		if err != nil {
			log.Println("client.Do faild,err:", err)
			continue
		} else {
			break
		}

	}

	defer resp.Body.Close()
	buf, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Println("ioutil.ReadAll failed,err:", err)
		return
	}
	fmt.Println(string(buf))
	return
}
