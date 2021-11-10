siusiu 会自动下载chromium，并通过 -c 参数指定路径

### 爬去深度为10，并输出到终端
crawlergo -m 10 url1 url2 url3... (must be same host)
### 输出格式为json,并输出到文件中
crawlergo -m 10 -o json --output-json xxx.json
### 自定义请求头
crawlergo -o json --custom-headers "{\"Cookie\": \"crawlergo=Cool\"}"


