package routers

import (
	"biu/controllers"
	"biu/py"

	"github.com/abiosoft/ishell"
)

//Init 初始化路由
func Init(shell *ishell.Shell) {
	//未找到命令时
	shell.NotFound(controllers.NotFoundHandler)
	//scan 端口扫描
	scanCmd := &ishell.Cmd{
		Name: "scan",
		Help: "主动扫描端口",
		Func: controllers.ScanController,
	}
	shell.AddCmd(scanCmd)
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
	guessCmd := &ishell.Cmd{
		Name: "guess",
		Help: "弱口令爆破器,支持:ssh,ftp,mysql,redis,mssql,postgresql,mongodb",
		Func: controllers.GuessController,
	}
	shell.AddCmd(guessCmd)
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
	baiduCmd := &ishell.Cmd{
		Name: "baidu",
		Help: "baidu url采集",
		Func: controllers.BaiduController,
	}
	shell.AddCmd(baiduCmd)
	//bing url 采集
	bingCmd := &ishell.Cmd{
		Name: "bing",
		Help: "bing url采集",
		Func: nil,
	}
	bingCmd.AddCmd(&ishell.Cmd{
		Name: "search",
		Help: "搜索关键字",
		Func: controllers.BingSearchController,
	})
	bingCmd.AddCmd(&ishell.Cmd{
		Name: "file",
		Help: "从文件中读取搜索关键字",
		Func: controllers.BingFileController,
	})
	// bingCmd.AddCmd(&ishell.Cmd{
	// 	Name: "sql-inject",
	// 	Help: "采集可能存在sql注入的url",
	// 	Func: controllers.BingSqlInjectController,
	// })
	shell.AddCmd(bingCmd)
	//google
	googleCmd := &ishell.Cmd{
		Name: "google",
		Help: "google url采集",
		Func: nil,
	}
	googleCmd.AddCmd(&ishell.Cmd{
		Name: "search",
		Help: "搜索关键字",
		Func: controllers.GoogleSearchController,
	})
	shell.AddCmd(googleCmd)
	//db
	dbCmd := &ishell.Cmd{
		Name: "db",
		Help: "操作数据库",
		Func: nil,
	}
	dbCmd.AddCmd(&ishell.Cmd{
		Name: "all",
		Help: "所有数据",
		Func: controllers.DbAllController,
	})
	dbCmd.AddCmd(&ishell.Cmd{
		Name: "urls",
		Help: "根据关键字获取采集到url",
		Func: controllers.DbURLsController,
	})
	shell.AddCmd(dbCmd)
	//whois
	whoisCmd := &ishell.Cmd{
		Name: "whois",
		Help: "whois查询",
		Func: controllers.WhoisController,
	}
	shell.AddCmd(whoisCmd)
	//目录相关的操作
	dirCmd := &ishell.Cmd{
		Name: "dir",
		Help: "目录相关的操作",
		Func: nil,
	}
	dirCmd.AddCmd(&ishell.Cmd{
		Name: "collect",
		Help: "采集某个项目的所有目录名",
		Func: controllers.DirCollectController,
	})
	shell.AddCmd(dirCmd)
	//py 脚本
	pyCmd := &ishell.Cmd{
		Name: "py",
		Help: "执行python脚本",
		Func: nil,
	}
	pyCmd.AddCmd(&ishell.Cmd{
		Name: "collect-proxy",
		Help: "代理采集",
		Func: func(c *ishell.Context) {
			py.Exec("collect-proxy.py", c.Args)
		},
	})
	pyCmd.AddCmd(&ishell.Cmd{
		Name: "backup-dict",
		Help: "生成网站备份字典",
		Func: func(c *ishell.Context) {
			py.Exec("backup-dict.py", c.Args)
		},
	})
	pyCmd.AddCmd(&ishell.Cmd{
		Name: "passwd-based-domain",
		Help: "基于域名生成若口令字典,常用于爆破网站后台密码",
		Func: func(c *ishell.Context) {
			py.Exec("passwd-based-domain.py", c.Args)
		},
	})
	pyCmd.AddCmd(&ishell.Cmd{
		Name: "passwd-based-userinfo",
		Help: "基于用户资料生成弱口令字典",
		Func: func(c *ishell.Context) {
			py.Exec("passwd-based-userinfo.py", c.Args)
		},
	})
	pyCmd.AddCmd(&ishell.Cmd{
		Name: "cms-fingerprint",
		Help: "cms指纹识别",
		Func: func(c *ishell.Context) {
			py.Exec("cms-fingerprint/main.py", c.Args)
		},
	})
	pyCmd.AddCmd(&ishell.Cmd{
		Name: "c-segment-scan",
		Help: "c段弱点发现",
		Func: func(c *ishell.Context) {
			py.BashExec("c-segment-scan/run.sh", c.Args)
		},
	})
	shell.AddCmd(pyCmd)
}
