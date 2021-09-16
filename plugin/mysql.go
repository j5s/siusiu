package plugin

import (
	"biu/models"
	"fmt"

	_ "github.com/go-sql-driver/mysql" //匿名导入 默认会自动执行该包中的init()方法
	"github.com/jinzhu/gorm"
)

//ConnectMySQLServer 连接MysQL服务端
//@param target 目标
//@param credential 凭据
func ConnectMySQLServer(target models.Target, credentail models.Credential) (result bool, err error) {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/mysql?charset=utf8mb4&parseTime=True",
		credentail.Username, credentail.Password, target.IP, target.Port)
	db, err := gorm.Open("mysql", dsn)
	if err != nil {
		return false, err
	}
	db.Close()
	return true, err
}
