package routers

import (
	"siusiu/controllers"
	"siusiu/pkg/exec"
	"siusiu/settings"

	"github.com/abiosoft/ishell"
)

//Init 初始化路由
func Init(shell *ishell.Shell) {
	//第三方工具
	for _, tool := range settings.AppConfig.Tools {
		shell.AddCmd(&ishell.Cmd{
			Name: tool["Name"],
			Help: tool["Help"],
			Func: func(c *ishell.Context) {
				exec.Bash(tool["Run"], c.Args)
			},
		})
	}
	//未找到命令时
	shell.NotFound(controllers.NotFoundHandler)
	//scan 端口扫描
	shell.AddCmd(&ishell.Cmd{
		Name: "port-scan",
		Help: "主动扫描端口",
		Func: controllers.ScanController,
	})
	//shodan 被动扫描
	shodanCmd := &ishell.Cmd{
		Name: "shodan",
		Help: "通过shodan被动扫描目标主机",
		Func: nil,
	}
	shodanCmd.AddCmd(&ishell.Cmd{
		Name: "credits",
		Help: "查询额度",
		Func: controllers.GetCreditsHandler,
	})
	shodanCmd.AddCmd(&ishell.Cmd{
		Name: "ports",
		Help: "被动端口扫描",
		Func: controllers.ShodanController,
	})
	shell.AddCmd(shodanCmd)
	//PasswdGuess 暴力破解
	shell.AddCmd(&ishell.Cmd{
		Name: "passwd-guess",
		Help: "弱口令爆破器,支持:ssh,ftp,mysql,redis,mssql,postgresql,mongodb",
		Func: controllers.GuessController,
	})
	//influxd 配置疏忽漏洞利用
	influxCmd := &ishell.Cmd{
		Name: "influx",
		Help: "influx 配置疏忽漏洞利用",
		Func: nil,
	}
	influxCmd.AddCmd(&ishell.Cmd{
		Name: "footprint",
		Help: "踩点,并尝试获取泄漏的敏感信息",
		Func: controllers.InfluxFootPrintController,
	})
	influxCmd.AddCmd(&ishell.Cmd{
		Name: "attack",
		Help: "利用jwt空秘钥漏洞",
		Func: controllers.InfluxAttackController,
	})
	shell.AddCmd(influxCmd)
	//baidu url 采集
	shell.AddCmd(&ishell.Cmd{
		Name: "baidu",
		Help: "baidu url采集",
		Func: func(c *ishell.Context) {
			exec.Python3("baidu.py", c.Args)
		},
	})
	//whois
	shell.AddCmd(&ishell.Cmd{
		Name: "whois",
		Help: "whois查询",
		Func: controllers.WhoisController,
	})
	//采集项目的目录名
	shell.AddCmd(&ishell.Cmd{
		Name: "dir-collector",
		Help: "采集某个项目的所有目录名",
		Func: controllers.DirCollectController,
	})
	shell.AddCmd(&ishell.Cmd{
		Name: "proxy-collector",
		Help: "代理采集",
		Func: func(c *ishell.Context) {
			exec.Python3("collect-proxy.py", c.Args)
		},
	})
	shell.AddCmd(&ishell.Cmd{
		Name: "backup-dict",
		Help: "生成网站备份字典",
		Func: func(c *ishell.Context) {
			exec.Python3("backup-dict.py", c.Args)
		},
	})
	shell.AddCmd(&ishell.Cmd{
		Name: "passwd-based-domain",
		Help: "基于域名生成若口令字典,常用于爆破网站后台密码",
		Func: func(c *ishell.Context) {
			exec.Python3("passwd-based-domain.py", c.Args)
		},
	})
	shell.AddCmd(&ishell.Cmd{
		Name: "passwd-based-userinfo",
		Help: "基于用户资料生成弱口令字典",
		Func: func(c *ishell.Context) {
			exec.Python3("passwd-based-userinfo.py", c.Args)
		},
	})
	shell.AddCmd(&ishell.Cmd{
		Name: "cms-fingerprint",
		Help: "cms指纹识别",
		Func: func(c *ishell.Context) {
			exec.Python3("cms-fingerprint/main.py", c.Args)
		},
	})
	shell.AddCmd(&ishell.Cmd{
		Name: "c-segment-scan",
		Help: "c段弱点发现",
		Func: func(c *ishell.Context) {
			exec.Python3("c-segment-scan/run.sh", c.Args)
		},
	})
}
