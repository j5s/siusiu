package exec

import (
	"context"
	"os"
	"os/exec"
	"os/signal"
	"siusiu/settings"
	"strings"

	"github.com/sirupsen/logrus"
)

//Bash bash 执行 shell脚本
func Bash(scriptName string, args []string) {
	execPath := settings.GetToolExecPath(scriptName)
	params := strings.Join(args, " ")
	CmdExec("/bin/bash", execPath, params)
}

//Python3 执行python3脚本
func Python3(scriptName string, args []string) {
	execPath := settings.GetToolExecPath("py/" + scriptName)
	params := strings.Join(args, " ")
	CmdExec("python3", execPath, params)
}

func CmdExec(name string, args ...string) {
	cmd := exec.Command(name, args...)
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt, os.Kill)
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	go func(ctx context.Context) {
	LOOP:
		for {
			select {
			case <-ctx.Done():
				break LOOP
			case <-c:
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
