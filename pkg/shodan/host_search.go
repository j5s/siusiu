package shodan

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"os"
)

//Host 主机信息
type Host struct {
	Ports []int `json:"ports"`
}

//ErrInvalidIP 非法IP
var ErrInvalidIP error = errors.New("非法的IP格式")

//SearchHostByIP 通过IP搜索主机信息
func (s *Client) SearchHostByIP(ip string) (*Host, error) {
	if ip == "" {
		return nil, ErrInvalidIP
	}
	resp, err := http.Get(fmt.Sprintf("%s/shodan/host/%s?key=%s", BaseURL, ip, s.apiKey))
	if err != nil {
		fmt.Fprintln(os.Stderr, "http.Get failed,err:", err)
		return nil, err
	}
	defer resp.Body.Close()
	var host Host
	err = json.NewDecoder(resp.Body).Decode(&host)
	if err != nil {
		fmt.Fprintln(os.Stderr, "json decode failed,err:", err)
		return nil, err
	}
	return &host, nil
}
