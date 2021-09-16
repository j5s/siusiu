package controllers

import (
	"biu/pkg/scanner"
	"biu/validate"

	"github.com/abiosoft/ishell"
	"github.com/sirupsen/logrus"
)

//ScanController 主动扫描
func ScanController(c *ishell.Context) {
	//1.解析主机
	hosts, err := validate.ReadIPList(c)
	if err != nil {
		logrus.Error("validate.ReadIPList(c) failed,err:", err)
		return
	}
	//2.解析端口范围
	ports, err := validate.ReadPortList(c)
	if err != nil {
		logrus.Error("validate.ReadPortList(c) failed,err:", err)
		return
	}
	//3.解析协程数
	count, err := validate.ReadGoRoutineCount(c)
	if err != nil {
		logrus.Error("validate.ReadPortList(c) failed,err:", err)
		return
	}
	//4.选择扫描方式
	mode := validate.ReadScanMode(c)
	//消费者
	portScanner := scanner.NewPortScanner(count, mode)
	portScanner.DealTask(count)
	//生产者:生产扫描任务
	portScanner.GenTask(hosts, ports)
	return
}
