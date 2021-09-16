package util

import (
	"fmt"
	"os"
)

//IsRoot 检查当前用户的uid是否等于0
func IsRoot() bool {
	return os.Geteuid() == 0
}

//CheckRoot 如果不是root，就退出
func CheckRoot() {
	if !IsRoot() {
		fmt.Fprintln(os.Stdout, "[*]程序运行需要root权限")
		os.Exit(0)
	}
}
