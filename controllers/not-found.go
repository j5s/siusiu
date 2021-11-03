package controllers

import (
	"log"
	"os"
	"os/exec"
	"runtime"
	"siusiu/settings"
	"strings"

	"github.com/abiosoft/ishell"
	"github.com/sirupsen/logrus"
)

//NotFoundHandler 未找到命令时处理函数
func NotFoundHandler(c *ishell.Context) {
	if c.Args[0] == "cd" {
		var dir string
		switch len(c.Args) {
		case 2:
			if strings.HasPrefix(c.Args[1], "$") {
				dir = os.Getenv(c.Args[1][1:])
			} else {
				dir = c.Args[1]
			}
		case 1:
			dir = os.Getenv("HOME")
		default:
			return
		}
		if err := os.Chdir(dir); err != nil {
			log.Println("os.Chdir failed,err:", err)
			return
		}
		c.SetPrompt(settings.GetShellPrompt())
		return
	}
	var cmd *exec.Cmd
	input := strings.Join(c.RawArgs, " ")

	if runtime.GOOS == "windows" {
		cmd = exec.Command("cmd", "/C", input) //windows
	} else {
		cmd = exec.Command("/bin/bash", "-c", input)
	}
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	if err := cmd.Run(); err != nil {
		logrus.Error("cmd.Run failed,err:", err)
		return
	}
}
