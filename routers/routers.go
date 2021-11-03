package routers

import (
	"io"
	"io/ioutil"
	"log"
	"os"
	"path"
	"siusiu/controllers"
	"siusiu/pkg/exec"
	"siusiu/settings"
	"strings"

	"github.com/abiosoft/ishell"
)

//Init 初始化路由
func Init(shell *ishell.Shell) error {
	//第三方工具
	for _, tool := range settings.AppConfig.Tools {
		shell.AddCmd(&ishell.Cmd{
			Name: tool["Name"],
			Help: tool["Help"],
			Func: func(run string) func(c *ishell.Context) {
				return func(c *ishell.Context) {
					list := []string{"|", ">", ">>", "<"}
					for i := range list {
						if strings.Contains(strings.Join(c.Args, " "), list[i]) {
							controllers.NotFoundHandler(c)
							return
						}
					}
					exec.Bash(run, c.Args)
				}
			}(tool["Run"]),
		})
	}
	demosCmd := &ishell.Cmd{
		Name: "demos",
		Help: "获取工具的使用样例",
		Func: nil,
	}
	demoesPath := path.Join(settings.AppConfig.MyVendorPath, "demos")
	markdowns, err := ioutil.ReadDir(demoesPath)
	if err != nil {
		log.Println("ioutil.ReadDir failed,err:", err)
		return err
	}
	for i := range markdowns {
		demosCmd.AddCmd(&ishell.Cmd{
			Name: markdowns[i].Name(),
			Help: markdowns[i].Name(),
			Func: func(markdown string) func(c *ishell.Context) {
				return func(c *ishell.Context) {
					filepath := path.Join(demoesPath, markdown)
					reader, err := os.Open(filepath)
					if err != nil {
						log.Println("os.Open failed,err:", err)
						return
					}
					if _, err := io.Copy(os.Stdout, reader); err != nil {
						log.Println("io.Copy failed,err:", err)
						return
					}
					return
				}
			}(markdowns[i].Name()),
		})
	}
	shell.AddCmd(demosCmd)
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
	return nil
}
