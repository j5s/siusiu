package controllers

import (
	"biu/pkg/scanner"
	"biu/validate"

	"github.com/abiosoft/ishell"
	"github.com/sirupsen/logrus"
)

//GuessController 爆破功能的处理函数
func GuessController(c *ishell.Context) {
	//1.读取主机端口文件
	targetList, err := validate.ReadTargetFilePath(c)
	if err != nil {
		logrus.Error("validate.ReadTargetFilePath(c) failed,err:", err)
		return
	}
	//2.读取用户名文件
	userList, err := validate.ReadUserDictFilePath(c)
	if err != nil {
		logrus.Error("validate.ReadUserDictFilePath(c) failed,err:", err)
		return
	}
	//3.读取密码文件
	passwdList, err := validate.ReadPasswdDictFilePath(c)
	if err != nil {
		logrus.Error("validate.ReadPasswdDictFilePath(c) failed,err:", err)
		return
	}
	//4.解析协程数
	count, err := validate.ReadGoRoutineCount(c)
	if err != nil {
		logrus.Error("validate.ReadPortList(c) failed,err:", err)
		return
	}
	//4.业务逻辑
	guesser := scanner.NewPasswdGuesser(passwdList, userList, targetList, count)
	//5.启动消费者
	guesser.PrintResult()
	//6.消费-生产者开始爆破
	guesser.DealTask()
	//7.生产者开始生产
	guesser.GenTask(passwdList, userList, targetList)
	//8.Summary
	guesser.Summary()
}
