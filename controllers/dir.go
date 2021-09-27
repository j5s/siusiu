package controllers

import (
	"biu/validate"
	"fmt"
	"io"
	"log"
	"os"
	"path/filepath"
	"time"

	"github.com/abiosoft/ishell"
	"github.com/sirupsen/logrus"
)

//DirCollectController 目录收集
func DirCollectController(c *ishell.Context) {

	path, err := validate.ReadDirPath(c)
	if err != nil {
		log.Println("validate.ReadDirPath failed,err:", err)
		return
	}
	filename := fmt.Sprintf("%v.log", time.Now().Unix())
	fp, err := os.Create(filename)
	if err != nil {
		logrus.Error("os.Create failed,err:", err)
		return
	}
	w := io.MultiWriter(os.Stdout, fp)
	if err := filepath.Walk(path,
		func(path string, info os.FileInfo, err error) error {
			if err != nil {
				return err
			}
			fmt.Fprintln(w, path)
			return nil
		}); err != nil {
		log.Println("filepath.Walk failed,err:", err)
		return
	}
	logrus.Infof("文件保存在:%s", filename)
}
