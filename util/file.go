package util

import (
	"biu/logger"
	"biu/models"
	"biu/plugin"
	"bufio"
	"errors"
	"os"
	"strconv"
	"strings"
)

//ErrFormat 格式错误
var ErrFormat error = errors.New("数据格式不正确,正确格式示例:110.220.142.148:22|ssh")

//PortProtocalMap 端口服务名常见的映射关系
var PortProtocalMap = map[int]string{
	21:    "FTP",
	22:    "SSH",
	3306:  "MYSQL",
	6379:  "REDIS",
	1433:  "MSSQL",
	5432:  "POSTGRESQL",
	27017: "MONGODB",
}

//ReadTargetList 读取目标列表
func ReadTargetList(filename string) ([]models.TargetWithHandler, error) {
	//1.添加文件描述符到进程的文件描述表
	file, err := os.Open(filename)
	if err != nil {
		logger.Log.Fatalf("Open ip List file err, %v", err)
	}
	defer file.Close()
	targets := make([]models.TargetWithHandler, 0, 1024)
	//2.逐行读取文件
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		if line == "" {
			continue
		}
		//0.拆分为ip和端口|协议两部分
		ipPort := strings.Split(line, ":")
		if len(ipPort) != 2 {
			logger.Log.Fatalf("strings.Split(line.':') failed")
		}
		ip := ipPort[0]
		portProtocal := strings.Split(ipPort[1], "|")
		//1.转化port为数字
		port, err := strconv.Atoi(portProtocal[0])
		if err != nil {
			logger.Log.Fatalf("strconv.Atoi(portProtocal[0]) failed")

		}
		//如果指明了服务类型
		var protocal string
		if len(portProtocal) == 2 {
			//2.转化协议为大写格式
			protocal = strings.ToUpper(portProtocal[1])

		} else {
			//如果没有指明服务类型
			var ok bool
			protocal, ok = PortProtocalMap[port]
			if !ok {
				logger.Log.Fatalf("请指定 %d 端口对应的协议类型", port)
			}
		}
		//3.添加到目标切片中去
		targets = append(targets, models.TargetWithHandler{
			Target: models.Target{
				IP:       ip,
				Port:     port,
				Protocal: protocal,
			},
			HandlerFunc: plugin.HandlerFuncMap[protocal],
		})

	}

	return targets, nil
}

//ReadList 读单字段文件
func ReadList(filename string) (list []string, err error) {
	file, err := os.Open(filename)
	if err != nil {
		logger.Log.Fatalf("Open password dict file err, %v", err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		if line != "" {
			list = append(list, line)
		}
	}
	//密码可能为空
	list = append(list, "")
	return list, err
}
