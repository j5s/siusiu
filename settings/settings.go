package settings

import (
	"fmt"
	"os"

	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"gopkg.in/fatih/color.v1"
)

//Config 配置
type Config struct {
	ShellPrompt  string              `mapstructure:"shell_prompt"`
	MyVendorPath string              `mapstructure:"my_vendor_path"`
	Tools        []map[string]string `mapstructure:"tools"`
}

//AppConfig App配置
var AppConfig *Config = new(Config)

//Init 初始化
func Init(filePath string) error {
	//指定配置文件
	viper.SetConfigFile(filePath)
	//读取配置文件
	if err := viper.ReadInConfig(); err != nil {
		fmt.Printf("viper.ReadInConfig() failed,err:%v\n", err)
		return err
	}
	//反序列化配置信息
	if err := viper.Unmarshal(AppConfig); err != nil {
		fmt.Printf("viper.Unmarshal(&Conf) failed,err:%v\n", err)
		return err
	}
	home := os.Getenv("HOME")
	AppConfig.MyVendorPath = fmt.Sprintf("%s/src/siusiu/myvendor", home)
	return nil
}

//GetShellPrompt 获取shell提示符
func GetShellPrompt() string {
	pwd, err := os.Getwd()
	if err != nil {
		logrus.Error("os.Getwd failed,err:", err)
		return fmt.Sprintf("%s > ", AppConfig.ShellPrompt)
	}
	return fmt.Sprintf("%s:%s > ", color.YellowString(AppConfig.ShellPrompt), color.GreenString(pwd))
}

//GetToolExecPath 获取工具执行路径
//@param scriptName 工具相对于myvendor的路径
func GetToolExecPath(scriptName string) string {
	return fmt.Sprintf("%s/%s", AppConfig.MyVendorPath, scriptName)
}
