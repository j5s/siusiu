host = input("请输入网站的域名(例如:www.baidu.com):")
host = host.split('.')
domain = '_'.join(host)
names = ['备份','backup','beifen','www','wwwroot',host[1],domain]
exts = ['.rar','.zip','.tar','.gz','.7z']
for name in names:
    for ext in exts:
        full_name = name+ext
        print(full_name)