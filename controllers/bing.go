package controllers

import (
	"biu/logic"
	"biu/validate"

	"github.com/abiosoft/ishell"
	"github.com/sirupsen/logrus"
)

//BingSearchController url采集
func BingSearchController(c *ishell.Context) {
	keyword, err := validate.ReadKeyword(c)
	if err != nil {
		logrus.Error("validate.ReadKeyword(c) failed,err:", err)
		return
	}
	routineCount, err := validate.ReadGoRoutineCount(c)
	if err != nil {
		logrus.Error("validate.ReadGoRoutineCount(c) failed,err:", err)
		return
	}
	logic.Search(keyword, routineCount)
}

func BingFileController(c *ishell.Context) {
	tails, err := validate.ReadFilePath(c)
	if err != nil {
		logrus.Error("validate.ReadFilePath(c) failed,err:", err)
		return
	}
	for line := range tails.Lines {
		go logic.Search(line.Text, 1)
	}
}
