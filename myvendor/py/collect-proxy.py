'''
代理采集程序
'''
import requests
import sys
import queue
import threading
from bs4 import BeautifulSoup as bs
import re
import pymysql
import time

headers = {
    "User-Agent": "Mozilla/5.0 (Macintosh; Intel Mac OS X 10.15; rv:84.0) Gecko/20100101 Firefox/84.0",
}




class Worker(threading.Thread):
    def __init__(self, inq, outq):
        threading.Thread.__init__(self)
        self._inq = inq
        self._outq = outq

    def run(self):
        while True:
            url = self._inq.get()
            try:
                self.fetcher(url)
            except:
                continue

    def fetcher(self, url):
        resp = requests.get(url=url, headers=headers)
        if(resp.status_code != 200):
            print("error status_code:%d" % (resp.status_code))
        tr_tags = re.findall(
            r"<tr><td>([^<]+)</td><td>(\d+)</td><td>([^<]+)</td><td>([^<]+)</td><td>[^<]+</td></tr>", resp.content.decode("gb2312"))
        for t in tr_tags:
            result = {
                "ip": t[0],
                "port": t[1],
                "area": t[2],
                "type": t[3]
            }
            self._outq.put(result)
            # print(result)


class Checker(threading.Thread):
    def __init__(self, inq):
        threading.Thread.__init__(self)
        self._inq = inq
        self._conn = pymysql.connect(host='localhost', user="root",passwd="root", db="hacktool")
        self._cur = self._conn.cursor()

    def run(self):
        while True:
            p = self._inq.get()
            ret_ip = self.checker(p)
            if(ret_ip is None):
                continue
            elif(ret_ip == p['ip']):
                print("get a useful proxy:%s:%s" % (p["ip"], p["port"]))
                self.saver(p)
            else:
                print("Useless proxy:%s:%s,because ret is %s" %
                      (p["ip"], p["port"], ret_ip))
        self._cur.close()
        self._conn.close()

    def checker(self,p):
        #查重
        sql = "select * from proxy where ip='%s' and port='%s'" % (p['ip'], p['port'])
        self._cur.execute(sql)
        row = self._cur.fetchall()
        if(len(row)>0):
            print("Duplicate proxy:%s:%s" % (p['ip'], p['port']))
            return None
        try:
            proxy = {"http": "http://"+p['ip']+":"+p['port']}
            resp = requests.get("http://icanhazip.com/",
                                proxies=proxy, headers=headers)
            if(resp.status_code!=200):
                print("error status code %d"%(resp.status_code))
                return None
            ret_ip = resp.text.strip()
            return ret_ip
        except:
            print("Useless proxy:%s:%s,because %s" %
                  (p["ip"], p["port"], "can't send request"))
            return None

    def saver(self, p):
        sql = "insert into proxy(ip,port,area,type,latest_check_time) values('%s','%s','%s','%s',%d)" % (
            p['ip'], p['port'], p['area'], p['type'], int(time.time()))
        self._cur.execute(sql)
        self._conn.commit()
        print("%s:%s插入成功" % (p['ip'], p['port']))


def main():
    inq = queue.Queue()
    outq = queue.Queue()
    workers = []
    workers_count = 10
    checkers = []
    checkers_count = 20
    for i in range(0, 2556):
        url = "http://www.66ip.cn/%d.html" % (i)
        inq.put(url)
    for i in range(workers_count):
        workers.append(Worker(inq, outq))
    for i in range(checkers_count):
        checkers.append(Checker(outq))
    for w in workers:
        w.start()
    for c in checkers:
        c.start()
    for c in checkers:
        c.join()
    for w in workers:
        w.join()


if __name__ == "__main__":
    main()
