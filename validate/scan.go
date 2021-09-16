package validate

import (
	"biu/util"
	"fmt"
	"net"

	"github.com/abiosoft/ishell"
	"github.com/sirupsen/logrus"
)

var (
	defaultPortRange    = "20-30,53,80-139,3000,3306,6379,8080,8086"
	defaultRoutineCount = "100"
	defaultMode         = "full"
)

//ReadIPList 读取IPList
func ReadIPList(c *ishell.Context) ([]net.IP, error) {
	c.Print("请输入要扫描的主机列表(支持:10.0.0.1 192.168.1.0/24 192.168.1.0-255 192.168.1.* 四种形式,多种格式之间可以用','连接):")
	iplist := c.ReadLine()
	hosts, err := util.GetIPList(iplist)
	if err != nil {
		logrus.Error("utils.GetIPList failed,err:", err)
		return nil, err
	}
	return hosts, nil
}

//ReadPortList 读取端口列表
func ReadPortList(c *ishell.Context) ([]int, error) {
	c.Print(fmt.Sprintf("请输入要扫描的端口范围(eg:21,22,80,20521-20530,默认值:%s):", defaultPortRange))
	portlist := c.ReadLine()
	if len(portlist) == 0 {
		portlist = defaultPortRange
	}
	ports, err := util.GetPorts(portlist)
	if err != nil {
		logrus.Error("util.GetPorts(portRange) failed,err:", err)
		return nil, err
	}
	return ports, err
}

//ReadScanMode 读取扫描方式
func ReadScanMode(c *ishell.Context) string {
	c.Print(fmt.Sprintf("请选择扫描方式(syn 半连接扫描,full 全连接扫描,默认为%s):", defaultMode))
	mode := c.ReadLine()
	switch mode {
	case "syn":
	case "full":
	default:
		mode = defaultMode
	}
	return mode
}
