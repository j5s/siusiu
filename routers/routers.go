package routers

import (
	"biu/controllers"

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
}

// var PasswdGuess = cli.Command{
// 	Name:        "guess",
// 	Usage:       "-i 主机文件名 -u 用户名字典 -p 密码字典 -r 协程数 -t 超时时间 -d debug 模式",
// 	Description: "弱口令爆破器,支持:ssh,ftp,mysql,redis,mssql,postgresql,mongodb",
// 	Action:      controllers.GuessController,
// 	Flags: []cli.Flag{
// 		cli.BoolFlag{Name: "debug, d", Usage: "debug模式"},
// 		cli.IntFlag{Name: "timeout, t", Value: 5, Usage: "连接的超时时间"},
// 		cli.IntFlag{Name: "routine_count, r", Value: 5000, Usage: "协程数量"},
// 		cli.StringFlag{Name: "ip_list, i", Value: "./dict/ip_list.txt", Usage: "主机列表文件的路径"},
// 		cli.StringFlag{Name: "user_dict, u", Value: "./dict/user.dic", Usage: "用户名字典文件的路径"},
// 		cli.StringFlag{Name: "pass_dict, p", Value: "./dict/pass.dic", Usage: "密码字典文件的路径"},
// 	},
// }
