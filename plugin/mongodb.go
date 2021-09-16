package plugin

import (
	"biu/models"
	"fmt"
	"time"

	"gopkg.in/mgo.v2"
)

//ConnectMongoDBServer 连接MongoDB服务端
func ConnectMongoDBServer(target models.Target, credential models.Credential) (result bool, err error) {
	url := fmt.Sprintf("mongodb://%v:%v@%v:%v/test",
		credential.Username, credential.Password, target.IP, target.Port)
	session, err := mgo.DialWithTimeout(url, time.Second)
	if err != nil {
		return false, err
	}
	defer session.Close()
	err = session.Ping()
	if err != nil {
		return false, err
	}
	return true, nil
}
