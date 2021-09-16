package plugin

import (
	"biu/models"
	"database/sql"
	"fmt"
)

//ConnectPostgreSQLServer 连接PostgreServer服务端
//@param target 目标
//@param credential 凭据
func ConnectPostgreSQLServer(target models.Target, credential models.Credential) (result bool, err error) {
	dataSourceName := fmt.Sprintf("postgres://%v:%v@%v:%v/postgres?sslmode=disable",
		credential.Username, credential.Password, target.IP, target.Port)
	db, err := sql.Open("postgres", dataSourceName)
	if err != nil {
		return false, err
	}
	err = db.Ping()
	if err != nil {
		return false, err
	}
	return true, err
}
