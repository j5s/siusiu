package controllers

import (
	"os"
	"siusiu/pkg/exec"
	"siusiu/settings"
	"strings"

	"github.com/abiosoft/ishell"
	"github.com/sirupsen/logrus"
)

//NotFoundHandler 未找到命令时处理函数
func NotFoundHandler(c *ishell.Context) {
	args := replaceVar(c.RawArgs)
	input := strings.Join(args, " ")
	if args[0] == "cd" {
		var dir string
		switch len(args) {
		case 1:
			dir = os.Getenv("HOME")
		default:
			return
		}
		if err := os.Chdir(dir); err != nil {
			logrus.Error("os.Chdir failed,err:", err)
			return
		}
		c.SetPrompt(settings.GetShellPrompt())
		return
	}
	exec.CmdExec("/bin/bash", "-c", input)
}

//replaceVar 替换命令中的变量
func replaceVar(rawArgs []string) []string {
	args := make([]string, 0, len(rawArgs))
	for _, arg := range rawArgs {
		//1.替换所有环境变量
		if strings.Contains(arg, "$") {
			arg = os.ExpandEnv(arg)
		}
		//2.替换所有内置命令
		for _, tool := range settings.AppConfig.Tools {
			if arg == tool["Name"] {
				arg = settings.GetToolExecPath(tool["Run"])
			}
		}
		args = append(args, arg)
	}
	return args
}
