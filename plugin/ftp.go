package plugin

import (
	"biu/logger"
	"biu/models"
	"fmt"
	"time"

	"github.com/jlaffaye/ftp"
)

//ConnectFTPServer 连接FTP服务器
//@param target 目标
//@param credential 凭据
func ConnectFTPServer(target models.Target, credential models.Credential) (result bool, err error) {
	//1.连接服务端
	conn, err := ftp.DialTimeout(fmt.Sprintf("%s:%d", target.IP, target.Port), 1*time.Second)
	if err != nil {
		logger.Log.Debugf("ftp.DialTimeout failed,err:%s", err.Error())
		return false, err
	}
	//2.连接成功,尝试登陆
	err = conn.Login(credential.Username, credential.Password)
	if err != nil {
		logger.Log.Debugf("conn.Login failed,err:%s", err.Error())
		return false, err
	}
	conn.Logout()
	return true, nil
}
