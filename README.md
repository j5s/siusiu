# siusiu （suitesuite）
一个用来管理suite 的suite，志在将渗透测试工程师从各种安全工具的学习和使用中解脱出来，减少渗透测试工程师花在安装工具、记忆工具使用方法上的时间和精力。

## Features

siusiu提供了一个shell控制台，通过该控制台，可以：

- 查看第三方安全工具列表
- 自动安装第三方安全工具
- 运行第三方安全工具
- 查看第三方安全工具的说明文档与使用样例（通过demos命令）

同时siusiu也支持非交互模式，便于siusiu被其他程序调用,例如:siusiu exec help

## Usage：

```
siusiu > help

Commands:
  GitHack                    .git泄漏利用脚本
  Glass                      针对资产列表的快速指纹识别工具
  SecList                    各种字典、webshell合集
  TPscan                     一键ThinkPHP漏洞检测
  Vulcan                     资产扫描工具(红队)
  XMLmining                  从xlsx、pptx、docx 文件的metadata中挖掘有用信息的工具
  arp-spoofing               局域网内主机扫描，ARP投毒、中间人攻击、敏感信息嗅探，HTTP报文嗅探
  backup-dict                生成网站备份字典
  baidu                      baidu url采集
  c-segment-scan             c段弱点发现
  clear                      clear the screen
  cms-fingerprint            cms指纹识别
  demos                      获取工具的使用样例
  dir-collector              采集某个项目的所有目录名
  dirsearch                  目录扫描器
  ds_store_exp               macOS .DS_Store文件泄漏利用脚本
  dvcs-ripper                SVN 泄漏利用脚本
  exit                       exit the program
  help                       display help
  influx                     influx 配置疏忽漏洞利用
  nmap                       端口扫描器
  one-for-all                一款功能强大的子域名收集工具
  passwd-based-domain        基于域名生成若口令字典,常用于爆破网站后台密码
  passwd-based-userinfo      基于用户资料生成弱口令字典
  passwd-guess               弱口令爆破器,支持:ssh,ftp,mysql,redis,mssql,postgresql,mongodb
  pocsuite3-cli              poc框架(命令行模式)
  pocsuite3-console          poc框架(控制台模式)
  port-scan                  主动扫描端口
  proxy-collector            代理采集
  shiro-attack               shiro反序列化漏洞综合利用工具(GUI)
  shodan                     通过shodan被动扫描目标主机
  sqlmap                     自动化sql注入工具
  url-collector              搜索引擎URL采集器(goole,bing)
  vim-swp-exp                vim swp 文件泄漏利用工具
  wafw00f                    waf指纹识别
  whois                      whois查询
  zenmap                     nmap-gui 版本,一个端口扫描器


siusiu >  
```

## Installation:

```
wget https://gitee.com/nothing-is-nothing/siusiu/raw/master/setup.sh
chmod +x setup.sh
./setup.sh
siusiu
```

## Screenshots

如果用户未安装pocsuite3，则自动下载 pocsuite3,然后自动运行 
![avatar](https://img-blog.csdnimg.cn/20211006160456729.png?x-oss-process=image/watermark,type_ZHJvaWRzYW5zZmFsbGJhY2s,shadow_50,text_Q1NETiBA5peg5Zyo5peg5LiN5Zyo,size_20,color_FFFFFF,t_70,g_se,x_16)

在siusiu控制台中运行sqlmap和dirsearch
![avatar](https://img-blog.csdnimg.cn/20211006160557298.png?x-oss-process=image/watermark,type_ZHJvaWRzYW5zZmFsbGJhY2s,shadow_50,text_Q1NETiBA5peg5Zyo5peg5LiN5Zyo,size_20,color_FFFFFF,t_70,g_se,x_16)


## Tested On  

- MacOS
- CentOS7
- Ubuntu

## Develop  

如果您有其他好的安全工具也想集成到siusiu中，可以按照如下步骤操作：  
step1.在siusiu安装目录（$HOME/src/siusiu）下创建对应的工具目录（建议以工具名命名,例如：dirsearch），并在该目录下创建该工具的shell脚本 run.sh，例如：

```shell
#!/bin/bash
base_path=$HOME/src
dirsearch_path=$base_path/dirsearch

function download {
    git clone https://github.com.cnpmjs.org/maurosoria/dirsearch.git $1
    cd $1
    pip3 install -r requirements.txt
}

#1.检查程序目录是否存在
if [ ! -d $dirsearch_path ]; then
    #2.如果不存在就下载
    echo "[*] download dirsearch..."
    download $dirsearch_path
fi
#运行dirsearch
python3 $dirsearch_path/dirsearch.py $*
```
step2. 在config.json 配置文件中添加对应工具，例如：
```
        {
            "Name": "dirsearch",
            "Help": "目录扫描器",
            "Run": "dirsearch/run.sh"
        },
```
其中name为工具名，help为工具描述，run为该工具的run.sh在myvendor目录下的相对路径

## 为工具编写demo文档

不知道你是否也曾有过这样的烦恼：每天疲于学习各种工具的使用方法,当真正需要使用某个工具的时候，却一时半会儿想不起某个工具怎么用，这时你翻开了你的笔记本，找呀找，终于找到了以前的笔记。  
关于这个问题，siusiu提供一种解决方案：将工具的使用文档或者常用demo集成在shell控制台中，需要时直接通过命令：demos+工具名 查看即可。  
你可以将你常用的一些命令demo，以markdown文档的方式写在 $HOME/src/siusiu/myvendor/demos 目录下，siusiu控制台会自动读取该目录。  
例如为sqlmap编写常用demo文档：

```markdown
# sqlmap demoes

```shell
# -m 批量扫描 —batch 全部采用默认行为，不向用户请求y/n,并且使用随机的user—agnet
sqlmap -m temp2.txt --batch --random-agent> result.txt

# 尝试获取所有数据库名
sqlmap -u url --dbs —-random-agent --batch

# 获取表名
sqlmap -u url --tables —-random-agent --batch

# 尝试获取所有用户:
sqlmap -u url --users --random-agent --batch

# 尝试获取账号密码:
sqlmap -u url --password --random-agent --batch

# 尝试获取当前用户:
sqlmap -u url --current-user --random-agent --batch

# 测试当前用户权限:
sqlmap -u url --is-dba --random-agent --batch

# 尝试写入木马,getshell
sqlmap -u url --os-shell --random-agent --batch

# 执行指定的sql语句
sqlmap.py -u url -v 1 --sql-query 'select top 20 * from City'
```

在siusiu控制台中通过 demos sqlmap.md 即可查看该文档。  