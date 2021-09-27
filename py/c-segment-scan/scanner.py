import threading
import requests
import sys
from IPy import IP
import queue
import re 

headers = {
    "UserAgent": "Mozilla/5.0 (Macintosh; Intel Mac OS X 10.15; rv:84.0) Gecko/20100101 Firefox/84.0"
}
class Worker(threading.Thread):
    def __init__(self,inq):
        threading.Thread.__init__(self)
        self._inq = inq
    def run(self):
        file = open("result.html","a+")
        while not self._inq.empty():
            target_url = self._inq.get()
            # print(target_url)
            try:
                resp = requests.get(url=target_url,timeout=6,headers=headers)
                if(resp.status_code == 200):
                    title = re.findall(r"<title>(.*?)</title>",resp.content.decode("utf-8"));
                    print("[*]Get! %s  %s"%(target_url,title[0]))
                    file.write("<a href='%s' target='_blank'>%s %s</a><br>"%(target_url,target_url,title[0]))
                    file.flush()
                else:
                    pass
                    # print("Error Status Code %d:%s"%(resp.status_code,url))
            except:
                # print("error")
                continue
        file.close()

def generateTaskes():
    q = queue.Queue()
    file = open("./py/c-segment-scan/nmap.xml",'r')
    content = file.read()
    ips = re.findall(r'<address addr="(.+?)" addrtype="ipv4"/>',content)

    # ips = IP("101.200.141.0/24")
    ports =['80','8080','3128','8081','9098','8888','443']
    protocals = ['http://']
    dirs = ["","/phpmyadmin","/admin","/login"]
    for ip in ips:
        for port in ports:
            for protocal in protocals:
                for dir in dirs:
                    target = protocal+str(ip)+":"+port+dir
                    q.put(target)
    for ip in ips:
        q.put("https://"+str(ip))
    return q


def main():
    q = generateTaskes()
    threads =[]
    threads_count = 200
    for i in range(threads_count):
        threads.append(Worker(q))
    for t in threads:
        t.start()
    for t in threads:
        t.join()

if __name__ == '__main__':
    main()
        
