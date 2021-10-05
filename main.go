package main

import (
	"siusiu/routers"

	"github.com/abiosoft/ishell"
	"github.com/abiosoft/readline"
)

func main() {
	shell := ishell.NewWithConfig(&readline.Config{
		Prompt: "siusiu > ",
	})
	routers.Init(shell)
	shell.Run()
}
