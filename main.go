package main

import (
	"siusiu/routers"
	"siusiu/settings"

	"github.com/abiosoft/ishell"
	"github.com/abiosoft/readline"
)

func main() {
	shell := ishell.NewWithConfig(&readline.Config{
		Prompt: settings.AppConfig.ShellPrompt,
	})
	routers.Init(shell)
	shell.Run()
}
