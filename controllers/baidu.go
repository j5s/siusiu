package controllers

import (
	"fmt"
	"io"
	"os"
	"os/exec"

	"github.com/abiosoft/ishell"
	"github.com/sirupsen/logrus"
)

func BaiduController(c *ishell.Context) {
	c.Print("请输入文件路径:")
	filepath := c.ReadLine()
	dstFile, err := os.OpenFile(filepath, os.O_RDWR|os.O_CREATE|os.O_APPEND, 777)
	if err != nil {
		logrus.Error("os.OpenFile failed,err:", err)
		return
	}
	w := io.MultiWriter(os.Stdout, dstFile)
	cmd := exec.Command("/Applications/MAMP/htdocs/security-tools/crawler/baidu/dist/baidu", "inurl:.php?id=")
	cmd.Stdin = os.Stdin
	cmd.Stderr = os.Stderr
	cmd.Stdout = w
	err = cmd.Run()
	if err != nil {
		fmt.Println("Execute Command failed:" + err.Error())
		return
	}

}
