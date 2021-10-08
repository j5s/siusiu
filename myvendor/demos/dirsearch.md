# dirsearch demos

-e 指定字典中%EXT% 的值，如果没有用-e 指定，字典中所有含有%EXT%的行都会被跳过
比如: index.%EXT% 指定 -e php,html,js 后会自动扫描 index index.php index.html index.js
如果已知某个站点不可能是 jsp 或者php的站，就可以用 -x 来排除字典中带有这类的后缀的项，提高扫描效率


dirsearch --random-agent --full-url -e php,html,js -u https://target 
dirsearch --random-agent --full-url -X jsp,php -u https://target 


如果是没有后缀的站，可以用--remove-extensions 移除字典中所有的后缀来进行扫描,这样字典更加有效，更小（大概可以减少3000行）扫描用时更短
dirsearch  --remove-extensions --full-url -u http://target

根据响应报文排除一些响应
dirsearch --exclude-texts="Not Found" --full-url --random-agent -u http://target

设置代理
dirsearch --exclude-texts="Not Found" --full-url --random-agent -u http://target --proxy http://127.0.0.1:8080 
dirsearch --exclude-texts="Not Found" --full-url --random-agent -u http://target --proxy-list proxy.txt

dirsearch 默认配置读取的 default.conf 文件，可以直接修改文件，将random-agent和full-url都开启。

指定字典:
dirsearch --exclude-texts="Not Found" --full-url --random-agent -u http://target --proxy http://127.0.0.1:8080 -w 字典路径

批量目录扫描:
比如说你有很多个目标站点需要目录扫描,你可以把他们保存在一个文件中,然后:
dirsearch --exclude-texts="Not Found" --full-url --random-agent -f xxx.txt --proxy http://127.0.0.1:8080 -o temp.txt


或者也可以从让dirsearch中标准输入中读
cat xxx.txt | grep "edu" | dirsearch --full-url --random-agent -o temp.txt --stdin


指定扫描某个子目录：
比如说你发现了某个站的后台/admin/ 就想要扫描一下/admin/这个子目录下有什么东西
dirsearch --full-url --random-agent --remove-extensions --exclude-texts="Not Found"  --subdirs admin/,uploads/ -u http://target


爆破配置文件:
dirsearch --full-url --random-agent --remove-extensions --suffixes ~ --prefixes . -u http://target 

这个条命令会将字典的大小扩大三倍，因为1行会衍生成3行：本行+带后缀的行+带前缀的行

爆破源码:
dirsearch --full-url --random-agent --remove-extensions -e zip --suffixes ~,.swp,.zip --prefixes . -u http://target


如果目标有限流或者为了反溯源，最好带上 --proxy-list 

为了提高扫描速度，可以使用method HEAD, 这样就减少的响应报文的大小。
dirsearch --full-url --random-agent --remove-extensions -e zip --suffixes ~,.swp,.zip --prefixes . -u http://target -m HEAD

如果返回429 表示 too many requests 被限流了。通常出现在你挂了某个代理的时候。不想看到429报告可以加 --skip-on-status 429
