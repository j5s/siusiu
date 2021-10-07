package main

import (
	"log"
	"siusiu/routers"
	"siusiu/settings"

	"github.com/abiosoft/ishell"
	"github.com/abiosoft/readline"
)

func main() {
	shell := ishell.NewWithConfig(&readline.Config{
		Prompt: settings.AppConfig.ShellPrompt,
	})
	if err := routers.Init(shell); err != nil {
		log.Println("routers.Init failed,err:", err)
		return
	}
	shell.Run()
}
