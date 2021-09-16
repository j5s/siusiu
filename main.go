package main

import (
	"biu/routers"

	"github.com/abiosoft/ishell"
	"github.com/abiosoft/readline"
)

func main() {
	shell := ishell.NewWithConfig(&readline.Config{
		Prompt: "biu > ",
	})
	routers.Init(shell)
	shell.Run()
}
