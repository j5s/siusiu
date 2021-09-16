package influxdb

import (
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"

	"github.com/sirupsen/logrus"
)

//CreateAdmin 创建管理员用户
func (c *Client) CreateAdmin(username string, password string) error {
	q := url.QueryEscape(fmt.Sprintf(`create user "%s" with password '%s' with all privileges`, username, password))
	url := fmt.Sprintf(`http://%s:%d/query?q=%s`, c.Host, c.Port, q)
	request, err := http.NewRequest(http.MethodPost, url, nil)
	if err != nil {
		logrus.Error("http.NewRequest failed,err:", err)
		return err
	}
	request.Header.Set("Authorization", fmt.Sprintf("Bearer %s", c.Token))
	resp, err := http.DefaultClient.Do(request)
	if err != nil {
		logrus.Error("http.DefaultClient.Do failed,err:", err)
		return err
	}
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}
	if resp.StatusCode == 200 {
		fmt.Printf("创建管理员级别用户成功,用户名:%s,密码:%s\n", username, password)
	} else {
		return errors.New("创建管理员用户失败,响应:" + string(body))
	}
	return nil
}
