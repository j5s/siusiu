package validate

import (
	"errors"
	"fmt"
	"os"
	"strconv"

	"github.com/abiosoft/ishell"
	"github.com/nxadm/tail"
	"github.com/sirupsen/logrus"
)

//ReadGoRoutineCount 读取协程数
func ReadGoRoutineCount(c *ishell.Context) (int, error) {
	c.Print(fmt.Sprintf("请输入协程数(默认为%s):", defaultRoutineCount))
	routineCount := c.ReadLine()
	if len(routineCount) == 0 {
		routineCount = defaultRoutineCount
	}
	r, err := strconv.Atoi(routineCount)
	if err != nil {
		logrus.Error("strconv.Atoi(routineCount) failed,err:", err)
		return r, err
	}
	return r, nil
}

//ReadKeyword 读取关键字
func ReadKeyword(c *ishell.Context) (string, error) {
	c.Print("keyword:")
	keyword := c.ReadLine()
	if len(keyword) == 0 {
		return "", errors.New("keyword 不能为空")
	}
	return keyword, nil
}

//ReadKey 读取秘钥
func ReadKey(c *ishell.Context, keyName string) (string, error) {
	key := os.Getenv("chinaz_key")
	if len(key) == 0 {
		return "", fmt.Errorf("环境变量%s为空,请在环境变量中写入对应的key", keyName)
	}
	return key, nil
}

//ReadDomain 读取域名
func ReadDomain(c *ishell.Context) (string, error) {
	c.Print("domain:")
	domain := c.ReadLine()
	if len(domain) == 0 {
		return "", errors.New("domain 不能为空")
	}
	return domain, nil
}

//ReadFilePath 读取关键字
func ReadFilePath(c *ishell.Context) (*tail.Tail, error) {
	c.Print("请输入文件路径:")
	filepath := c.ReadLine()
	if len(filepath) == 0 {
		return nil, errors.New("filepath 不能为空")
	}
	tails, err := tail.TailFile(filepath, tail.Config{
		ReOpen:    false,
		Follow:    false,
		MustExist: true,
		Poll:      true,
	})
	if err != nil {
		fmt.Println("tail file failed,err:", err)
		return nil, err
	}
	return tails, nil
}
