# siusiu （suitesuite）
一个用来管理suite 的suite

## Features

siusiu提供了一个shell控制台，通过该控制台，可以：

- 查看第三方安全工具列表
- 自动安装第三方安全工具
- 运行第三方安全工具

## Usage：
```
siusiu > help

Commands:
  GitHack                    .git泄漏利用脚本
  Glass                      针对资产列表的快速指纹识别工具
  SecList                    各种字典、webshell合集
  XMLmining                  从xlsx、pptx、docx 文件的metadata中挖掘有用信息的工具
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
  passwd-based-domain        基于域名生成若口令字典,常用于爆破网站后台密码
  passwd-based-userinfo      基于用户资料生成弱口令字典
  passwd-guess               弱口令爆破器,支持:ssh,ftp,mysql,redis,mssql,postgresql,mongodb
  pocsuite3-cli              poc框架(命令行模式)
  pocsuite3-console          poc框架(控制台模式)
  port-scan                  主动扫描端口
  proxy-collector            代理采集
  shodan                     通过shodan被动扫描目标主机
  sqlmap                     自动化sql注入工具
  url-collector              搜索引擎URL采集器(goole,bing)
  vim-swp-exp                vim swp 文件泄漏利用工具
  whois                      whois查询
  zenmap                     nmap-gui 版本,一个端口扫描器
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
step1 在siusiu安装目录（$HOME/src/siusiu）下的创建对应的工具目录（建议以工具名命名），并在该目录下创建该工具的自动下载和自动运行脚本 run.sh
run.sh demo:

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