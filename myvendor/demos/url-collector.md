#### 根据指定关键字采集
url-collector -k ".php?id=" 

#### 批量采集文件中所有的关键字，并将结果保存到result.txt
url-collector -f google-dork.txt -o result.txt

#### 和sqlmap联动
url-collector -f google-dork.txt -o result.txt
sqlmap -m result.txt --batch --random-agents

#### 默认采用google镜像站点，如果你可以访问外网，可以手动指定搜索引擎为google
url-collector -e google -k ".php?id="

#### 将常用配置写到配置文件中
url-collector -c config.json

#### 基于搜索引擎的子域名收集
url-collector -k "site:qq.com" -f domain

#### 搜索可能存在sql注入漏洞的url
url-collector -k "inurl:.asp?id=" 
url-collector -k "inurl:'.php?id='+inurl:.cn"

### 搜索thinkphp的站点
url-collector -k "intext:十年磨一剑 - 为API开发设计的高性能框架" -f protocol_domain
