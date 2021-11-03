package controllers

import (
	"context"
	"fmt"
	"log"
	"os"
	"os/exec"
	"os/signal"
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
	fmt.Println("intput:", input)
	if runtime.GOOS == "windows" {
		cmd = exec.Command("cmd", "/C", input) //windows
	} else {
		cmd = exec.Command("/bin/bash", "-c", input)
	}
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	ch := make(chan os.Signal, 1)
	signal.Notify(ch, os.Interrupt, os.Kill)
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	go func(ctx context.Context) {
	LOOP:
		for {
			select {
			case <-ctx.Done():
				break LOOP
			case <-ch:
				cmd.Process.Release()
				cmd.Process.Kill()
				break LOOP
			}
		}
	}(ctx)
	if err := cmd.Run(); err != nil {
		logrus.Error("cmd.Run failed,err:", err)
		return
	}
}
