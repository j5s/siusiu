package util

import (
	"errors"
	"fmt"
	"net"
	"strconv"
	"strings"

	"github.com/malfunkt/iprange"
)

var ErrInvalidPortRange error = errors.New("非法的端口范围")

//GetIPList 解析多IP:string表示的ip 转化为 []net.IP 支持:10.0.0.1 192.168.1.0/24 192.168.1.0-255 192.168.1.* 四种形式
func GetIPList(ipStr string) ([]net.IP, error) {
	addressList, err := iprange.ParseList(ipStr)
	if err != nil {
		return nil, err
	}
	list := addressList.Expand()
	return list, err
}

//GetPorts 解析多端口
func GetPorts(portsStr string) ([]int, error) {
	ports := make([]int, 0, 1024)
	if "" == portsStr {
		return ports, nil
	}
	//"22,0-100,8080" => ["22","0-100","8080"]
	portStrList := strings.Split(portsStr, ",")
	for _, portStr := range portStrList {
		portStr = strings.TrimSpace(portStr)
		//1.如果含有 - ,表示端口范围
		if strings.Contains(portStr, "-") {
			var minPort, maxPort int
			_, err := fmt.Sscanf(portStr, "%d-%d", &minPort, &maxPort)
			if err != nil {
				return nil, err
			}
			if minPort < 0 || maxPort > 65535 || minPort > maxPort {
				return nil, ErrInvalidPortRange
			}
			for i := minPort; i <= maxPort; i++ {
				ports = append(ports, i)
			}
			//如果不含 - ,直接添加到端口数组中
		} else {
			port, err := strconv.Atoi(portStr)
			if err != nil {
				return nil, err
			}
			ports = append(ports, port)
		}
	}
	return ports, nil
}
