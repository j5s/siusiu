package shodan

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
)

//APIInfo 当前账户API的使用情况
type APIInfo struct {
	QueryCredits int `json:"query_credits"` //可用的查询额度
	ScanCredits  int `json:"scan_credits"`  //可用的扫描额度
}

//GetAPIInfo 请求当前账户的API使用情况
func (s *Client) GetAPIInfo() (*APIInfo, error) {
	//1.发送请求
	resp, err := http.Get(fmt.Sprintf("%s/api-info?key=%s", BaseURL, s.apiKey))
	if err != nil {
		fmt.Fprintln(os.Stderr, "http.Get failed,err:", err)
		return nil, err
	}
	defer resp.Body.Close()
	//2.反序列化json
	var ret APIInfo
	err = json.NewDecoder(resp.Body).Decode(&ret)
	if err != nil {
		fmt.Fprintln(os.Stderr, "json decode failed,err:", err)
		return nil, err
	}
	return &ret, nil
}
