package plugin

import (
	"fmt"
	"siusiu/models"

	"github.com/go-redis/redis"
)

//ConnectRedisServer 连接Redis服务端
//@target 目标
//@credential 登录凭据（用户名+密码）
func ConnectRedisServer(target models.Target, credentail models.Credential) (result bool, err error) {
	option := &redis.Options{
		Addr:     fmt.Sprintf("%s:%d", target.IP, target.Port),
		Password: credentail.Password, // redis只有密码没有用户名
		DB:       0,                   // use default DB
	}
	rdb := redis.NewClient(option)
	defer rdb.Close()
	_, err = rdb.Ping().Result()
	if err != nil {
		return false, err
	}
	return true, err
}
