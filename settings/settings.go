package settings

import (
	"fmt"
	"os"
)

//Config 配置
type Config struct {
	ShellPrompt  string
	MyVendorPath string
	Tools        []map[string]string
}

func newAppConfig() *Config {
	home := os.Getenv("HOME")
	tools := []map[string]string{
		{
			"Name": "sqlmap",
			"Help": "自动化sql注入工具",
			"Run":  "sqlmap/run.sh",
		},
		{
			"Name": "pocsuite3-cli",
			"Help": "poc框架(命令行模式)",
			"Run":  "pocsuite3/run-cli.sh",
		},
		{
			"Name": "pocsuite3-console",
			"Help": "poc框架(控制台模式)",
			"Run":  "pocsuite3/run-console.sh",
		},
		{
			"Name": "dirsearch",
			"Help": "目录扫描器",
			"Run":  "dirsearch/run.sh",
		},
		{
			"Name": "url-collector",
			"Help": "搜索引擎URL采集器(goole,bing)",
			"Run":  "url-collector/run.sh",
		},
		{
			"Name": "zenmap",
			"Help": "nmap-gui 版本,一个端口扫描器",
			"Run":  "nmap/gui/run.sh",
		},
		{
			"Name": "nmap",
			"Help": "端口扫描器",
			"Run":  "nmap/cli/run.sh",
		},
		{
			"Name": "GitHack",
			"Help": ".git泄漏利用脚本",
			"Run":  "GitHack/run.sh",
		},
	}

	return &Config{
		ShellPrompt:  "siusiu > ",
		MyVendorPath: fmt.Sprintf("%s/src/siusiu/myvendor", home),
		Tools:        tools,
	}
}

//AppConfig App配置
var AppConfig *Config = newAppConfig()
