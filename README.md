# biu 网安工具箱

## Usage：
```
Commands:
  backup-dict                生成网站备份字典
  baidu                      baidu url采集
  c-segment-scan             c段弱点发现
  clear                      clear the screen
  cms-fingerprint            cms指纹识别
  dir-collector              采集某个项目的所有目录名
  dirsearch                  目录扫描器
  exit                       exit the program
  help                       display help
  influx                     influx 配置疏忽漏洞利用
  nmap                       端口扫描器
  passwd-based-domain        基于域名生成若口令字典,常用于爆破网站后台密码
  passwd-based-userinfo      基于用户资料生成弱口令字典
  passwd-guess               弱口令爆破器,支持:ssh,ftp,mysql,redis,mssql,postgresql,mongodb
  port-scan                  主动扫描端口
  proxy-collector            代理采集
  shodan                     通过shodan被动扫描目标主机
  url-collector              搜索引擎URL采集器
  whois                      whois查询
  zenmap                     nmap-gui 版本,一个端口扫描器
```
## 使用方法:

```
go run main.go 或 直接运行编译好的可执行文件
```

## 截图

![在这里插入图片描述](https://img-blog.csdnimg.cn/806838be04a24ca4b8e89203384f4842.png?x-oss-process=image/watermark,type_ZHJvaWRzYW5zZmFsbGJhY2s,shadow_50,text_Q1NETiBA5peg5Zyo5peg5LiN5Zyo,size_20,color_FFFFFF,t_70,g_se,x_16)


phpmywind 指纹识别思路：

fofa 搜索4g.php

识别： /install/
http://gdzqhy.com/install/
如果包含关键字 phpmywind 就是
