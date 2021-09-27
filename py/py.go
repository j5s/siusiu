package py

import (
	"fmt"
	"io"
	"log"
	"os"
	"os/exec"
	"os/signal"
	"strings"
	"time"

	"github.com/sirupsen/logrus"
)

const pyPath = "/Users/mac/Desktop/Go/biu/py"
const outputPath = "/Users/mac/Desktop/Go/biu/output"

//Exec 执行py脚本
func Exec(scriptName string, args []string) {
	cmdExec("python3", scriptName, args)
}

//BashExec bash 执行 shell脚本
func BashExec(scriptName string, args []string) {
	cmdExec("/bin/bash", scriptName, args)
}

func cmdExec(command, scriptName string, args []string) {
	filename := strings.ReplaceAll(scriptName, "/", "-")
	filepath := fmt.Sprintf("%s/%s-%v.log", outputPath, filename, time.Now().Unix())
	defer logrus.Info("[*] file saved at:", filepath)
	dstFile, err := os.OpenFile(filepath, os.O_RDWR|os.O_CREATE|os.O_APPEND, 777)
	if err != nil {
		logrus.Error("os.OpenFile failed,err:", err)
		return
	}
	w := io.MultiWriter(os.Stdout, dstFile)

	cmd := exec.Command(command, fmt.Sprintf("%s/%s", pyPath, scriptName), strings.Join(args, " "))
	cmd.Stdin = os.Stdin
	cmd.Stdout = w
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
