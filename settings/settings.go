package settings

//Config 配置
type Config struct {
	ShellPrompt  string
	MyVendorPath string
}

//AppConfig App配置
var AppConfig Config = Config{
	ShellPrompt:  "siusiu >",
	MyVendorPath: "$HOME/src/siusiu/myvendor",
}
