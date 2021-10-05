'''
百度url采集程序
'''
import requests
import sys
import queue
import threading
from bs4 import BeautifulSoup as bs
import re
import optparse

headers = {
    "User-Agent": "Mozilla/5.0 (Macintosh; Intel Mac OS X 10.15; rv:84.0) Gecko/20100101 Firefox/84.0",
}


class BaiduURLGenerator(threading.Thread):
    """ baidu url 生成器 """
    def __init__(self, out_chan,keyword="",filepath=""):
        threading.Thread.__init__(self)
        self._out_chan = out_chan
        self._keyword=keyword
        self._filepath=filepath
    
    def run(self):
        if(len(self._keyword)>0):
            self.get_baidu_url(self._keyword)
        elif(len(self._filepath)>0):
            with open(filepath, 'r') as f: 
                while True:
                    keyword=f.readline()
                    if(len(keyword)==0):
                        break
                    self.get_baidu_url(keyword)
    
    def get_baidu_url(self,keyword):
        """根据关键词向管道中输送baidu url"""
        for i in range(0, 760,10):
            url = "https://www.baidu.com/s?wd=%s&pn=%d"%(keyword,i)
            self._out_chan.put(url)
    
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
        a_tags = soup.find_all(
            name="a", attrs={"class": None, "data-click": re.compile(r".+")})
        for a in a_tags:
            fake_url = a['href']
            resp = requests.get(url=fake_url, headers=headers, timeout=8)
            if(resp.status_code == 200):
                print(resp.url)



        





def main(keyword="",filepath=""):
    q = queue.Queue()
    threads = []
    threads_count = 10
    #消费者：从管道中拿 baidu url
    for i in range(threads_count):
        threads.append(Worker(q))
    for w in threads:
        w.start()
    #生产者：向管道中发送baidu url
    url_gen = BaiduURLGenerator(q,keyword,filepath)
    url_gen.start()
    #等待线程任务完成
    url_gen.join()
    for w in threads:
        w.join()
 


if __name__ == "__main__":
    flag=optparse.OptionParser(description="Baidu Search")
    flag.add_option("-k","--keyword",dest="keyword",help="keyword e.g: inurl:?id=")
    flag.add_option("-f","--filepath",dest="filepath",help="file path,e.g:urls.txt")
    (options,args)=flag.parse_args()
   
    if(options.keyword and len(options.keyword.strip())>0):
        keyword=options.keyword.strip()
        print("searching:%s...."%keyword)
        main(keyword)
    elif(options.filepath and len(options.filepath.strip())>0):
        filepath=options.filepath.strip()
        main(filepath=filepath)
    else:
        flag.print_help()
    
