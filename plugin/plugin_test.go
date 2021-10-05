package plugin

import (
	"siusiu/models"
	"testing"
)

func TestConnectSSHServer(t *testing.T) {
	items := []struct {
		credentials models.Credential
		target      models.Target
		result      bool
	}{
		{
			models.Credential{"user", "user"},
			models.Target{"192.168.1.100", 22, "SSH"},
			false,
		},
		{
			models.Credential{"root", "123456"},
			models.Target{"101.200.142.148", 22, "SSH"},
			true,
		},
		{
			models.Credential{"root", "X8sherlock1895"},
			models.Target{"101.200.142.148", 22, "SSH"},
			false,
		},
	}
	for _, item := range items {
		result, err := ConnectSSHServer(item.target, item.credentials)
		if err != nil {
			t.Error(err, result)
		}
		if result != item.result {
			t.Errorf("%v,%v,expect:%v,got:%v", item.target, item.credentials, item.result, result)
		}
	}
}

func TestConnectMySQLServer(t *testing.T) {
	items := []struct {
		credentials models.Credential
		target      models.Target
		result      bool
	}{
		{
			models.Credential{"root", "root"},
			models.Target{"127.0.0.1", 3306, "MYSQL"},
			true,
		},
		{
			models.Credential{"root", "123456"},
			models.Target{"127.0.0.1", 3306, "MYSQL"},
			false,
		},
		{
			models.Credential{"root", ""},
			models.Target{"127.0.0.1", 3306, "MYSQL"},
			false,
		},
	}
	for _, item := range items {
		result, _ := ConnectMySQLServer(item.target, item.credentials)
		if result != item.result {
			t.Errorf("%v,%v,expect:%v,got:%v", item.target, item.credentials, item.result, result)
		}
	}
}

func TestConnectRedisServer(t *testing.T) {
	items := []struct {
		credentials models.Credential
		target      models.Target
		result      bool
	}{
		{
			models.Credential{"", "root"},
			models.Target{"127.0.0.1", 6379, "REDIS"},
			true,
		},
		{
			models.Credential{"", "123456"},
			models.Target{"127.0.0.1", 6379, "REDIS"},
			true,
		},
		{
			models.Credential{"", ""},
			models.Target{"127.0.0.1", 6379, "REDIS"},
			false,
		},
	}
	for _, item := range items {
		result, err := ConnectRedisServer(item.target, item.credentials)
		if err != nil {
			t.Error(err)
		}
		if result != item.result {
			t.Errorf("%v,%v,expect:%v,got:%v", item.target, item.credentials, item.result, result)
		}
	}
}

func TestConnectFTPServer(t *testing.T) {
	items := []struct {
		credentials models.Credential
		target      models.Target
		result      bool
	}{
		{
			models.Credential{"admin", "tGAA1-TsAuHzdA29"},
			models.Target{"127.0.0.1", 21, "FTP"},
			true,
		},
		{
			models.Credential{"admin", "Wwrp6kiEgY3LTeIC"},
			models.Target{"127.0.0.1", 21, "FTP"},
			false,
		},
		{
			models.Credential{"admin", "admin123"},
			models.Target{"127.0.0.1", 21, "FTP"},
			true,
		},
	}
	for _, item := range items {
		result, err := ConnectFTPServer(item.target, item.credentials)
		if err != nil {
			t.Error(err)
		}
		if result != item.result {
			t.Errorf("%v,%v,expect:%v,got:%v", item.target, item.credentials, item.result, result)
		}
	}
}
