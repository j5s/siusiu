package validate

import (
	"siusiu/models"
	"siusiu/util"

	"github.com/abiosoft/ishell"
)

//ReadTargetFilePath 读取目标
func ReadTargetFilePath(c *ishell.Context) ([]models.TargetWithHandler, error) {
	c.Print("请输入要爆破的目标文件的路径:")
	iplist := c.ReadLine()
	return util.ReadTargetList(iplist)
}

//ReadUserDictFilePath 读取用户名文件
func ReadUserDictFilePath(c *ishell.Context) (list []string, err error) {
	c.Print("请输入用户名字典文件的路径:")
	userDict := c.ReadLine()
	return util.ReadList(userDict)
}

//ReadPasswdDictFilePath 读取密码文件
func ReadPasswdDictFilePath(c *ishell.Context) (list []string, err error) {
	c.Print("请输入密码字典文件的路径:")
	passDict := c.ReadLine()
	return util.ReadList(passDict)
}
