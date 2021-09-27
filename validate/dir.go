package validate

import (
	"errors"

	"github.com/abiosoft/ishell"
)

//ReadDirPath 读取目录的位置
func ReadDirPath(c *ishell.Context) (string, error) {
	return Read(c, "请输入目录的绝对路径:")
}

//Read 读取变量
func Read(c *ishell.Context, tip string) (string, error) {
	c.Print(tip)
	l := c.ReadLine()
	if len(l) == 0 {
		return "", errors.New("输入值不能为空")
	}
	return l, nil
}
