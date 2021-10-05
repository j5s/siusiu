package exec

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"os/signal"
	"strings"
)

const vendorPath = "/Users/mac/bin/biu/myvendor"

//Bash bash 执行 shell脚本
func Bash(scriptName string, args []string) {
	cmdExec("/bin/bash", scriptName, args)
}

//Python3 执行python3脚本
func Python3(scriptName string, args []string) {
	cmdExec("python3", "py/"+scriptName, args)
}

func cmdExec(command, scriptName string, args []string) {
	cmd := exec.Command(command, fmt.Sprintf("%s/%s", vendorPath, scriptName), strings.Join(args, " "))
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	go func() {
		c := make(chan os.Signal, 1)
		signal.Notify(c, os.Interrupt, os.Kill)
		_ = <-c
		cmd.Process.Kill()
	}()
	if err := cmd.Run(); err != nil {
		log.Println("cmd.Run failed,err:", err)
		return
	}
}
