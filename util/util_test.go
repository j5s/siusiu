package util

import (
	"net"
	"reflect"
	"testing"
)

func TestGetIPList(t *testing.T) {
	items := []struct {
		IPStr string
		IP    []net.IP
	}{
		{"192.168.1.0/24", []net.IP{}},
		{"192.168.1.*", []net.IP{}},
		{"192.168.1.0-100", []net.IP{}},
		{"192.168.1.0,192.168.1.1", []net.IP{}},
	}
	for _, item := range items {
		ip, err := GetIPList(item.IPStr)
		if err != nil {
			t.Errorf("getIPlist failed,err:%v", err)
			continue
		}
		if false == reflect.DeepEqual(ip, item.IP) {
			t.Errorf("GetIPList(%s);got %v;expected:%v", item.IPStr, ip, item.IP)
		}

	}

}

func TestGetPorts(t *testing.T) {
	items := []struct {
		PortStr string
		Ports   []int
	}{
		{"22,23,8080", []int{}},
		{"20521-20530,80,22,23,21", []int{}},
		{"100-1", []int{}},
		{"dsad-dasd", []int{}},
		{"", []int{}},
		{"dasd", []int{}},
	}
	for _, item := range items {
		ports, err := GetPorts(item.PortStr)
		if err != nil {
			t.Errorf("getPorts failed,err:%v", err)
			continue
		}
		if false == reflect.DeepEqual(ports, item.Ports) {
			t.Errorf("GetPorts(%s);got %v;expected:%v", item.PortStr, ports, item.Ports)
		}
	}
}
