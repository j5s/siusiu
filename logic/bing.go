package logic

import (
	"biu/dao/bolt"
	"biu/vars"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"strings"
)

//Search 搜索
func Search(keyword string, routineCount int) {
	URLCh := make(chan string, 1024)
	//消费者
	for i := 0; i < routineCount; i++ {
		go save(fetch(URLCh), keyword)
	}
	//生产者
	searchURL := fmt.Sprintf("%s/search?q=%s&ensearch=1", vars.BingURL, url.QueryEscape(keyword))
	URLCh <- searchURL
}

//save 保存数据到es
func save(resultCh <-chan string, keyword string) {
	for result := range resultCh {
		if strings.Contains(result, "gov.cn") {
			continue
		}
		bolt.AddSearchResult(keyword, result)
	}
}

//fetch 从搜索页面中提取文件的URL以及下页的URL
func fetch(searchItemCh chan string) <-chan string {
	resultURLCh := make(chan string, 1024)
	go func() {
		for searchItem := range searchItemCh {
			//1.发送请求
			request, err := http.NewRequest(http.MethodGet, searchItem, nil)
			if err != nil {
				log.Println("http.NewRequest failed,err:", err)
				continue
			}
			request.Header.Set("User-Agent", "Mozilla/5.0 (iPhone; CPU iPhone OS 11_0 like Mac OS X) AppleWebKit/604.1.38 (KHTML, like Gecko) Version/11.0 Mobile/15A372 Safari/604.1")
			resp, err := http.DefaultClient.Do(request)
			if err != nil {
				log.Printf("http.Get(%s) failed,err:%v", searchItem, err)
				continue
			}
			defer resp.Body.Close()
			//2.解析响应 寻找指定文件的URL
			buf, err := ioutil.ReadAll(resp.Body)
			matches := vars.AtagRe.FindAllStringSubmatch(string(buf), -1)
			for _, match := range matches {
				resultURLCh <- match[1]
			}
			//3.寻找“下一页URL”
			matches = vars.NextPageRe.FindAllStringSubmatch(string(buf), -1)
			for _, match := range matches {
				searchItemCh <- vars.BingURL + match[1]
			}
		}

	}()
	return resultURLCh
}
