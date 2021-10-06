package settings

import (
	"fmt"
	"os"
)

//Config 配置
type Config struct {
	ShellPrompt  string
	MyVendorPath string
}

func newAppConfig() *Config {
	home := os.Getenv("HOME")
	return &Config{
		ShellPrompt:  "siusiu > ",
		MyVendorPath: fmt.Sprintf("%s/src/siusiu/myvendor", home),
	}
}

//AppConfig App配置
var AppConfig *Config = newAppConfig()
