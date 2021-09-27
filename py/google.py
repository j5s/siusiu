import requests
import random
from lxml import etree

open('url.txt', 'w')
url = "https://g.luciaz.me/search?q=inurl:.php?id=%s 公司&amp;start=%s"  # 这里是谷歌语法
head = {
    'User-Agent': 'Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/92.0.4515.159 Safari/537.36'
}
password = random.randint(50, 200)
print("当前随机数为：%d" % password)
print("自动采集id参数为 %d 的公司URL" % password)
for i in range(20):  # 默认采集20个，有需求自己改
    newurl = format(url % (password, (i+1)))
    shuju = requests.get(url=newurl, headers=head).text
    data = etree.HTML(shuju)
    li_list = data.xpath('//div[@class="eqAnXb"]/div/div/div/div[1]/div/div/div/a/@href')[0]
    with open("url.txt", "a", encoding="utf-8") as ww:
        ww.write(li_list+"\n")
    print("当前已完成：%s" % i)
