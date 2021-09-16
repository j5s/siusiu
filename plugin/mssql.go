package plugin

import (
	"biu/models"
	"database/sql"
	"fmt"
)

//ConnectMSSQLServer 连接MSSQL服务端
func ConnectMSSQLServer(target models.Target, credential models.Credential) (result bool, err error) {
	dataSourceName := fmt.Sprintf("server=%v,port=%v;user id=%v;password=%v;database=master",
		target.IP, target.Port, credential.Username, credential.Password)
	db, err := sql.Open("mssql", dataSourceName)
	if err != nil {
		return false, err
	}
	err = db.Ping()
	if err != nil {
		return false, err
	}
	return true, nil
}
