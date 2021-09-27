package controllers

import (
	"biu/dao/bolt"
	"biu/validate"
	"fmt"
	"io"
	"os"
	"time"

	"github.com/abiosoft/ishell"
	"github.com/sirupsen/logrus"
)

func DbURLsController(c *ishell.Context) {
	keyword, err := validate.ReadKeyword(c)
	if err != nil {
		logrus.Error("validate.ReadKeyword(c) failed,err:", err)
		return
	}
	results, err := bolt.GetSearchResult(keyword)
	if err != nil {
		logrus.Error("bolt.GetSearchResult failed,err:", err)
		return
	}
	filename := fmt.Sprintf("%v.log", time.Now().Unix())
	fp, err := os.Create(filename)
	if err != nil {
		logrus.Error("os.Create failed,err:", err)
		return
	}

	w := io.MultiWriter(os.Stdout, fp)
	for i := range results {
		fmt.Fprintln(w, results[i].URL)
	}
	c.Println("文件保存在:", filename)
}

//DbAllController 显示所有数据
func DbAllController(c *ishell.Context) {
	results, err := bolt.GetAll()
	if err != nil {
		logrus.Error("bolt.GetAll failed,err:", err)
		return
	}
	filename := fmt.Sprintf("%v.log", time.Now().Unix())
	fp, err := os.Create(filename)
	if err != nil {
		logrus.Error("os.Create failed,err:", err)
		return
	}

	w := io.MultiWriter(os.Stdout, fp)
	for i := range results {
		fmt.Fprintln(w, results[i].URL)
	}
	c.Println("文件保存在:", filename)
}
