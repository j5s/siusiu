package plugin

import (
	"fmt"
	"net"
	"siusiu/logger"
	"time"

	"siusiu/models"

	"golang.org/x/crypto/ssh"
)

//ConnectSSHServer 连接SSH服务端
//@target 目标
//@credential 登录凭据（用户名+密码）
func ConnectSSHServer(target models.Target, credential models.Credential) (result bool, err error) {
	//1.客户端配置
	config := &ssh.ClientConfig{
		User: credential.Username,
		Auth: []ssh.AuthMethod{
			ssh.Password(credential.Password),
		},
		Timeout: time.Second,
		HostKeyCallback: func(hostname string, remote net.Addr, key ssh.PublicKey) error {
			return nil
		},
	}
	//2.连接服务端
	client, err := ssh.Dial("tcp", fmt.Sprintf("%s:%d", target.IP, target.Port), config)
	if err != nil {
		logger.Log.Debug("ssh.Dial failed,err:", err)
		return false, err
	}
	//3.连接成功
	defer client.Close()
	//4.测试是否可以执行命令
	session, err := client.NewSession()
	if err != nil {
		logger.Log.Debug("client.NewSession failed,err:", err)
		return true, err
	}
	err = session.Run("ls")
	if err != nil {
		logger.Log.Debug("session.Run(ls) failed,err:", err)
		return true, err
	}
	return true, nil
}
