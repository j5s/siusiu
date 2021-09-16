package controllers

import (
	"biu/pkg/shodan"
	"biu/util"
	"fmt"
	"os"

	"github.com/abiosoft/ishell"
	"github.com/sirupsen/logrus"
)

var shodanKey string = os.Getenv("shodan_key")

//GetCreditsHandler 查询额度
func GetCreditsHandler(c *ishell.Context) {
	if len(shodanKey) == 0 {
		c.Println("请在环境变量中写入您的shodan key 或者现在输入:")
		shodanKey = c.ReadLine()
	}
	c.Println("正在查询额度中...")
	client := shodan.NewClient(shodanKey)
	info, err := client.GetAPIInfo()
	if err != nil {
		fmt.Fprintln(os.Stderr, "client.GetAPIInfo failed,err:", err)
		return
	}
	fmt.Printf("查询额度:%d\n扫描额度:%d\n", info.QueryCredits, info.ScanCredits)
}

//ShodanController shodan被动扫描
func ShodanController(c *ishell.Context) {
	//1.解析ip地址
	c.Print("请输入要扫描的主机列表(支持:10.0.0.1 192.168.1.0/24 192.168.1.0-255 192.168.1.* 四种形式,多种格式之间可以用','连接):")
	iplist := c.ReadLine()
	hosts, err := util.GetIPList(iplist)
	if err != nil {
		logrus.Error("utils.GetIPList failed,err:", err)
		return
	}
	//2.被动扫描
	if len(shodanKey) == 0 {
		c.Println("请在环境变量中写入您的shodan key 或者现在输入:")
		shodanKey = c.ReadLine()
	}
	client := shodan.NewClient(shodanKey)
	for _, host := range hosts {
		result, err := client.SearchHostByIP(host.String())
		if err != nil {
			logrus.Error("client.SearchHostByIP(host.String()) failed,err:", err)
			return
		}
		for _, port := range result.Ports {
			fmt.Printf("%s:%d\n", host.String(), port)
		}
		if len(result.Ports) != 0 {
			fmt.Println()
		}
	}
	return
}
