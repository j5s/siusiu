package controllers

import (
	"log"
	"os"
	"os/exec"
	"runtime"
	"strings"

	"github.com/abiosoft/ishell"
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
		return
	}
	var cmd *exec.Cmd
	input := strings.Join(c.Args, " ")
	if runtime.GOOS == "windows" {
		cmd = exec.Command("cmd", "/C", input) //windows
	} else {
		cmd = exec.Command("/bin/bash", "-c", input)
	}
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	if err := cmd.Run(); err != nil {
		c.Println("cmd.Run failed,err:", err)
		return
	}
}
