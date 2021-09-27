import threading
import requests
import queue
import hashlib
import re
class Worker(threading.Thread):
    def __init__(self, q, r, url):
        threading.Thread.__init__(self)
        self.q = q
        self.r = r
        self.size = q.qsize()
        self.url = url

    def run(self):
        while not self.q.empty():
            item = self.q.get()
            url = self.url+item['url']

            try:
                resp = requests.get(url, timeout=1)
                if(resp.status_code == 200):
                    # print("%s\tlength:%d" % (url, len(resp.text)))
                    if(item['md5'] != ""):
                        cms = self.checkMd5(resp.content, item)
                        if(cms != None):
                            self.r.put(cms)
                    if(item["re"] != ""):
                        cms = self.checkReg(resp.text, item)
                        if(cms != None):
                            self.r.put(cms)
                else:
                    pass
                    # print("%s\terror status_code:%s" % (url, resp.status_code))
            except Exception as e:
                # print(e)
                pass

    def checkMd5(self, content, item):
        md5 = hashlib.md5(content)
        if(md5 == item['md5']):
            return item["name"]
        else:
            return None

    def checkReg(self, content, item):
        if re.search(item['re'], content):
            return item["name"]
        else:
            return None