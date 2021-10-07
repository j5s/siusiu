# dirsearch demos

```shell

dirsearch -e php,html,js -u https://target
# -e 指定字典中%EXT% 的值，如果没有用-e 指定，字典中所有含有%EXT%的行都会被跳过
# 比如: index.%EXT% 指定 -e php,html,js 后会自动扫描 index index.php index.html index.js
# 如果已知某个站点不可能是 jsp 或者php的站，就可以用 -x 来排除字典中带有这类的后缀的项，提高扫描效率

dirsearch -x jsp,php -u https://target
#如果是没有后缀的站，可以用--remove-extensions 移除字典中所有的后缀来进行扫描

dirsearch -u https://target --remove-extensions

#如果想根据响应字符排除某些
```