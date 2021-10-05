package controllers

import (
	"siusiu/models"
	"siusiu/validate"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/abiosoft/ishell"
	"github.com/sirupsen/logrus"
)

func WhoisController(c *ishell.Context) {
	domain, err := validate.ReadDomain(c)
	if err != nil {
		logrus.Error("validate.ReadDomain(c) failed,err:", err)
		return
	}
	key, err := validate.ReadKey(c, "chinaz_key")
	if err != nil {
		logrus.Error("validate.ReadKey(c) failed,err:", err)
		return
	}
	api := fmt.Sprintf("https://apidatav2.chinaz.com/single/newicp?key=%s&domain=%s", key, domain)
	fmt.Println("正在查询ICP域名备案信息:", api)
	resp, err := http.Get(api)
	if err != nil {
		logrus.Error("http.Get(api) failed,err:", err)
		return
	}
	defer resp.Body.Close()
	buf, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		logrus.Error("ioutil.ReadAll failed,err:", err)
		return
	}
	var r models.Resp
	if err := json.Unmarshal(buf, &r); err != nil {
		log.Println("json.Unmarshal failed,err:", err)
		return
	}
	fmt.Println("公司名称:", r.Result.CompanyName)
	fmt.Println("公司类型:", r.Result.CompanyType)
	fmt.Println("网站名称:", r.Result.SiteName)
	fmt.Println("网站主页:", r.Result.MainPage)
	fmt.Println("网站备案号:", r.Result.SiteLicense)
	fmt.Println("网站所有者:", r.Result.Owner)
}
