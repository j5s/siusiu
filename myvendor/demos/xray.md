```shell
## 单点扫描（不爬去网页内容）
xray webscan -u url
xray webscan -url-file file 
## 查看所有poc
xray webscan -l
## 使用指定的poc
xray webscan -p 指定poc -u url
xray webscan -u http://127.0.0.1:8080  -p poc-yaml-thinkphp5023-method-rce --html-output result.html
xray webscan -p "poc-yaml-thinkphp-v6-file-write,poc-yaml-thinkphp5-controller-rce,poc-yaml-thinkphp5023-method-rce" -u url --html-output result.html
xray webscan --plugins thinkphp,xss -u url
## 让xray成为浏览器和服务器的中间人
### 在当前目录下生成ca证书，获取客户端的信任(记得将证书导入火狐浏览器)
xray genca
### 监听端口(该命令会读取当前目录下生成的ca.crt和ca.key文件,并自动生成配置文件config.yaml,修改hostname_allowed可以限定扫描范围)
xray webscan --listen 127.0.0.1:7777 --html-output result.html
xray webscan --plugins cmd-injection,sqldet --listen 127.0.0.1:7777
## 爬虫模式(可以在配置文件中添加cookie，实现登录后的爬虫扫描)
xray webscan --basic-crawler http://testphp.vulnweb.com/ --html-output result.html
## 服务扫描
xray servicescan --target 127.0.0.1:8009
xray servicescan --target-file 1.file 
## 官方文档
https://docs.xray.cool
```
