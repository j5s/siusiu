<!--
 * @Date: 2021-01-18 23:53:20
 * @LastEditors: 无在无不在
 * @LastEditTime: 2021-01-19 00:14:50
-->
## C段web服务扫描器的意义：
+ 信息收集：找到目标系统在C段的其他服务
+ 寻找短板: 找到目标系统的其他弱点
+ 资产识别: 发现C段内的资产信息

e.g:
子域名扫描：

test.demo.com
damin.demo.com

设计点：
1. 指定端口 
2. 线程数可控
3. 对前端进行设计：数据可视化 或者 数据挖掘


nmap -sn -PE --min-hostgroup 1024 --min-parallelism 1024 -oX nmap.xml 101.200.142.0/24

-sn 不进行端口扫描，只进行ping检测
-PE 通过ICMP echo来判定主机是否存活
--min-hostgroup 1024 最小分组设置为1024个IP地址
--min-parallelism 1024 将探针的数量设置最小为1024（1024个线程）