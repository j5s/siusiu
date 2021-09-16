package models

import "fmt"

//Target 目标
type Target struct {
	IP       string
	Port     int
	Protocal string
}

//Serialize 序列化
func (t *Target) Serialize() string {
	return fmt.Sprintf("%s:%d|%s", t.IP, t.Port, t.Protocal)
}

//ServiceHandlerFunc 服务的处理函数
type ServiceHandlerFunc func(target Target, credential Credential) (result bool, err error)

//TargetWithHandler 带有处理函数的目标
type TargetWithHandler struct {
	Target
	HandlerFunc ServiceHandlerFunc
}

//Credential 登录凭证
type Credential struct {
	Username string
	Password string
}

//Serialize 序列化
func (c *Credential) Serialize() string {
	return fmt.Sprintf("username:%s password:%s", c.Username, c.Password)
}

//Task 一个爆破任务
type Task struct {
	TargetWithHandler
	Credential
}

//Check 执行爆破任务
func (t *Task) Check() (result bool, err error) {
	return t.TargetWithHandler.HandlerFunc(t.Target, t.Credential)
}

//Serialize 序列化
func (t *Task) Serialize() string {
	return fmt.Sprintf("%s username:%s password:%s", t.Target.Serialize(), t.Username, t.Password)
}
