package exec

import (
	"context"
	"fmt"
	"log"
	"os"
	"os/exec"
	"os/signal"
	"siusiu/settings"
	"strings"
)

//Bash bash 执行 shell脚本
func Bash(scriptName string, args []string) {
	cmdExec("/bin/bash", scriptName, args)
}

//Python3 执行python3脚本
func Python3(scriptName string, args []string) {
	cmdExec("python3", "py/"+scriptName, args)
}

func cmdExec(command, scriptName string, args []string) {
	cmd := exec.Command(command, fmt.Sprintf("%s/%s", settings.AppConfig.MyVendorPath, scriptName), strings.Join(args, " "))
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
		log.Println("cmd.Run failed,err:", err)
		return
	}
}
