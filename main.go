package main

import (
	"biu/dao/bolt"
	"biu/routers"

	"github.com/abiosoft/ishell"
	"github.com/abiosoft/readline"
	"github.com/sirupsen/logrus"
)

func main() {
	shell := ishell.NewWithConfig(&readline.Config{
		Prompt: "biu > ",
	})
	if err := bolt.Init(); err != nil {
		logrus.Error("bolt.Init failed,err:", err)
		return
	}
	defer bolt.Close()
	routers.Init(shell)
	shell.Run()
}
