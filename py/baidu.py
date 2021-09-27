'''
百度url采集程序
'''
import requests
import sys
import queue
import threading
from bs4 import BeautifulSoup as bs
import re

headers = {
    "User-Agent": "Mozilla/5.0 (Macintosh; Intel Mac OS X 10.15; rv:84.0) Gecko/20100101 Firefox/84.0",   
}
class Worker(threading.Thread):
    def __init__(self, q):
        threading.Thread.__init__(self)
        self._q = q

    def run(self):
        while True:
            url = self._q.get()
            try:
                self.fetcher(url)
            except:
                continue
         

    def fetcher(self, url):
        resp = requests.get(url=url, headers=headers)
        if(resp.status_code != 200):
            print("error status_code:%d" % (resp.status_code))
        soup = bs(resp.content, "lxml")
        a_tags = soup.find_all(name="a", attrs={"class":None,"data-click":re.compile(r".+")})
        for a in a_tags:
            fake_url = a['href']
            resp = requests.get(url=fake_url,headers=headers,timeout=8)
            if(resp.status_code == 200):
                print(resp.url)
                path_arr = resp.url.split("/")
                index = path_arr[0]+path_arr[1]+path_arr[2]
                print("首页地址:%s"%(index))
                
        

def main(keyword):
    q = queue.Queue()
    threads = []
    threads_count = 10
    for i in range(0, 760,10):
        url = "https://www.baidu.com/s?wd=%s&pn=%d"%(keyword,i)
        q.put(url)
    for i in range(threads_count):
        threads.append(Worker(q))
    for w in threads:
        w.start()
    for w in threads:
        w.join()


if __name__ == "__main__":
    if len(sys.argv) != 2:
        print("必须带一个参数:搜索关键词 eg:%s keyword"%(sys.argv[0]))
        sys.exit()
    else:
        main(sys.argv[1])
    