package plugin

import "siusiu/models"

//HandlerFuncMap 服务和处理函数之间的映射关系
var HandlerFuncMap map[string]models.ServiceHandlerFunc

func init() {
	HandlerFuncMap = make(map[string]models.ServiceHandlerFunc, 7)
	HandlerFuncMap["FTP"] = ConnectFTPServer
	HandlerFuncMap["SSH"] = ConnectSSHServer
	HandlerFuncMap["MYSQL"] = ConnectMySQLServer
	HandlerFuncMap["MSSQL"] = ConnectMSSQLServer
	HandlerFuncMap["YONGODB"] = ConnectMongoDBServer
	HandlerFuncMap["REDIS"] = ConnectRedisServer
	HandlerFuncMap["POSTGRESQL"] = ConnectPostgreSQLServer
}
