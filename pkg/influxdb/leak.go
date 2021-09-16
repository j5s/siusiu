package influxdb

import (
	"biu/models"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"

	"github.com/mitchellh/mapstructure"
	"github.com/sirupsen/logrus"
)

//GetDebugVars 请求debug/vars
func (c *Client) GetDebugVars() (leakinfo models.LeakInfo, err error) {
	fmt.Println("正在测试是否存在敏感信息泄漏...")
	url := fmt.Sprintf("http://%s:%d/debug/vars", c.Host, c.Port)
	resp, err := http.Get(url)
	if err != nil {
		logrus.Error("http.Get failed,err:", err)
		return leakinfo, err
	}
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		logrus.Error("ioutil.ReadAll failed,err:", err)
		return leakinfo, err
	}

	if strings.Contains(string(body), "authentication") {
		fmt.Println("目标主机开启了pprof-enabled...")
		fmt.Println(string(body))
		return leakinfo, errors.New("目标主机开启了pprof-auth-enabled")
	}
	fmt.Println("目标主机未开启pprof-auth-enabled,正在尝试解析...")

	var info map[string]interface{}
	if err := json.Unmarshal(body, &info); err != nil {
		logrus.Error("json.Unmarshal failed,err:", err)
		return leakinfo, err
	}
	if err := json.Unmarshal(body, &leakinfo); err != nil {
		logrus.Error("json.Unmarshal failed,err:", err)
		return leakinfo, err
	}
	databases := make([]models.DataBaseInfo, 0)
	for k, v := range info {

		if strings.HasPrefix(k, "database:") == false {
			continue
		}
		database := new(models.DataBaseInfo)
		if err := mapstructure.Decode(v, database); err != nil {
			logrus.Error("mapstructrue.Decode failed,err:", err)
			return leakinfo, err
		}
		databases = append(databases, *database)
	}
	leakinfo.Databases = databases
	return leakinfo, nil
}

//Ping 获取版本，测试是否开启了
func (c *Client) Ping() (version *models.Version, err error) {
	url := fmt.Sprintf("http://%s:%d/ping", c.Host, c.Port)
	resp, err := http.Get(url)
	if err != nil {
		logrus.Error("http.Get failed,err:", err)
		return nil, err
	}
	for k, v := range resp.Header {
		if k == "X-Influxdb-Version" {
			return models.NewVersion(v[0])
		}
	}
	return nil, errors.New("X-Influxdb-Version not found")
}

//IsAuthEnable 探测是否开启了身份认证
//auth-enabled = true
func (c *Client) IsAuthEnable() (bool, error) {
	q := url.QueryEscape("show users")
	url := fmt.Sprintf("http://%s:%d/query?q=%s", c.Host, c.Port, q)
	resp, err := http.Get(url)
	if err != nil {
		logrus.Error("http.Get failed,err:", err)
		return true, err
	}
	if resp.StatusCode == 400 || resp.StatusCode == 401 {
		return true, nil
	}
	return false, nil
}

//Follow 跟踪http客户端对 /write 和query的请求
func (c *Client) Follow() error {
	fmt.Println("正在跟踪其他http客户端的write和query请求...")
	url := fmt.Sprintf("http://%s:%d/debug/requests", c.Host, c.Port)
	resp, err := http.Get(url)
	if err != nil {
		logrus.Error("http.Get failed,err:", err)
		return err
	}
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}
	if len(string(body)) == 0 {
		fmt.Println("目前没有其他客户端向服务端发送请求")
	} else {
		fmt.Println(string(body))
	}
	return nil
}
