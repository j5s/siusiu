"""
seebug poc 采集程序
"""
import requests
import threading
from bs4 import BeautifulSoup as bs
import queue
import json
import time
"""
	//下载地址 /vuldb/downloadPoc/15335
	//兑换poc POST /vuldb/exchange/98358
	// {"type":"poc","anonymous":false}
	//兑换响应
	// {"status":false,"message":"\u60a8\u6240\u62e5\u6709\u7684 KB \u4e0d\u8db3"
"""

"""
思路：
1.访问所有page获取url
2.模拟兑换操作,保存下载地址
"""
headers = {
    "Host": "www.seebug.org",
    "User-Agent": "Mozilla/5.0 (Macintosh; Intel Mac OS X 10.15; rv:84.0) Gecko/20100101 Firefox/84.0",
    "Accept": "text/html,application/xhtml+xml,application/xml;q=0.9,image/webp,*/*;q=0.8",
    "Accept-Language": "zh-CN,zh;q=0.8,zh-TW;q=0.7,zh-HK;q=0.5,en-US;q=0.3,en;q=0.2",
    "Accept-Encoding": "gzip, deflate",
    "Connection": "close",
    "Referer": "https://www.seebug.org/vuldb/vulnerabilities?category=&order_time=1&order_rank=1&has_all=default&has_vm=default&submitTime=all&has_affect=default&has_poc=true&has_detail=default&level=all&page=1",
    "Cookie": '__jsluid_s=42e6cbb95503752296aa2163de2f4c05; Hm_lvt_6b15558d6e6f640af728f65c4a5bf687=1610768386,1610768457,1610768472,1610769176; csrftoken=jJ83hvOcUKCEZywXf1dFrD6Dd8h1N8GP; Hm_lpvt_6b15558d6e6f640af728f65c4a5bf687=1610792252; sessionid=fzrc2m8wvf9074x3t7ziua33tv08xx5u; messages="5940c806d556d063f4617688b18f0bbc6de07916$[[\"__json_message\"\0540\05425\054\"Login succeeded. Welcome\054 274d155d.\"]]',
    "Upgrade-Insecure-Requests": "1"
}
post_headers = {
    "Host": "www.seebug.org",
    "User-Agent": "Mozilla/5.0 (Macintosh; Intel Mac OS X 10.15; rv:84.0) Gecko/20100101 Firefox/84.0",
    "Accept": "application/json, text/javascript, */*; q=0.01",
    "Accept-Language": "zh-CN,zh;q=0.8,zh-TW;q=0.7,zh-HK;q=0.5,en-US;q=0.3,en;q=0.2",
    "Accept-Encoding": "gzip, deflate",
    "Content-Type": "application/x-www-form-urlencoded; charset=UTF-8",
    "X-CSRFToken": "jJ83hvOcUKCEZywXf1dFrD6Dd8h1N8GP",
    "X-Requested-With": "XMLHttpRequest",
    "Content-Length": "32",
    "Origin": "https://www.seebug.org",
    "Connection": "close",
    "Referer": "https://www.seebug.org/",
    "Cookie": '__jsluid_s=42e6cbb95503752296aa2163de2f4c05; Hm_lvt_6b15558d6e6f640af728f65c4a5bf687=1610768386,1610768457,1610768472,1610769176; csrftoken=jJ83hvOcUKCEZywXf1dFrD6Dd8h1N8GP; Hm_lpvt_6b15558d6e6f640af728f65c4a5bf687=1610795779; sessionid=fzrc2m8wvf9074x3t7ziua33tv08xx5u; messages="5940c806d556d063f4617688b18f0bbc6de07916$[[\"__json_message\"\0540\05425\054\"Login succeeded. Welcome\054 274d155d.\"]]"'
}

params = {
    "type": "poc",
    "anonymous": "false"
}


class Worker(threading.Thread):
    def __init__(self, q):
        threading.Thread.__init__(self)
        self._q = q

    def run(self):
        while True:
            url = self._q.get()
            self.fetcher(url)

    def fetcher(self, url):
        resp = requests.get(url=url, headers=headers)
        if(resp.status_code != 200):
            print("error status_code:%d" % (resp.status_code))
        soup = bs(resp.content, "lxml")
        vulids = soup.find_all(name="a", attrs={"class": "vul-title"})
        for v in vulids:
            title = v['title']
            id = v['href'].split('-')[1]
            exchange_url = "https://www.seebug.org/vuldb/exchange/"+str(id)
            print(exchange_url)
            data = json.dumps(params)
            referer = "https://www.seebug.org/vuldb/ssvid-"+str(id)
            post_headers["Referer"] = referer
            time.sleep(0.5)
            resp = requests.post(
                url=exchange_url, headers=post_headers, data=data)
            print(resp.status_code)
            if(resp.status_code == 200):
                j = json.loads(resp.content)
                print(j['message'])
                print(type(j['status']))
            # 下载poc
            if(resp.status_code == 403 or j['status'] == True):
                download_url = "https://www.seebug.org/vuldb/downloadPoc/"+id
                print("downloading...:%s"%(download_url))
                referer = "https://www.seebug.org/vuldb/ssvid-"+str(id)
                headers["Referer"] = referer
                print("downloading referer is %s"%(headers["Referer"]))
                resp = requests.get(url=download_url, headers=headers)
                print("downloading status code:%d"%(resp.status_code))
                if(resp.status_code == 200):
                    filename = "ssvid-"+str(id)+".txt"
                    print("writing %s" % (filename))
                    f = open("./poc/"+filename, 'w')
                    f.write(resp.text)
                    f.close()


def main():
    q = queue.Queue()
    threads = []
    threads_count = 10
    for i in range(200, 205):
        url = "https://www.seebug.org/vuldb/vulnerabilities?category=&order_time=1&order_rank=1&has_all=default&has_vm=default&submitTime=all&has_affect=default&has_poc=true&has_detail=default&level=all&page=" + \
            str(i)
        q.put(url)
    for i in range(threads_count):
        threads.append(Worker(q))
    for w in threads:
        w.start()
    for w in threads:
        w.join()


if __name__ == "__main__":
    main()
