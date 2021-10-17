package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"siusiu/routers"
	"siusiu/settings"

	"github.com/abiosoft/ishell"
	"github.com/abiosoft/readline"
)

func main() {
	//0.接收命令行参数
	var configFilePath string
	flag.StringVar(&configFilePath, "config", fmt.Sprintf("%s/src/siusiu/config.json", os.Getenv("HOME")), "指定配置文件的路径")
	flag.Parse()
	//1.加载配置
	if err := settings.Init(configFilePath); err != nil {
		fmt.Printf("init setting failed,err:%v\n", err)
		return
	}
	//2.初始化shell
	shell := ishell.NewWithConfig(&readline.Config{
		Prompt: settings.AppConfig.ShellPrompt,
	})
	//3.初始化路由
	if err := routers.Init(shell); err != nil {
		log.Println("routers.Init failed,err:", err)
		return
	}
	shell.Run()
}
