import queue
import builtwith
import sys
from cmslist import cmsdata
from worker import Worker
class ScanCms():
    def __init__(self, url,thread_count):
        self.url = url
        self.thread_count = thread_count
        self.result = {}
        pass

    def run(self):
        # 调用builtwith识别web开发信息
        result = self.built()
        print(result)
        # 调用识别方法，识别web的cms信息
        cms = self.cms()
        result['cms'] = cms
        print()
        print(result)

    def built(self):
        try:
            result = builtwith.parse(self.url)
        except:
            result = None
        if result:
            return result
        else:
            return {}

    def cms(self):
        threads = []
        q = queue.Queue()
        r = queue.Queue()
        for item in cmsdata:
            q.put(item)
        size = q.qsize()
        for i in range(self.thread_count):
            threads.append(Worker(q, r, self.url))
        for t in threads:
            t.start()
        while r.empty():
            current = q.qsize()
            percent = ((size-current)/size)*100
            msg  ="\r[*]剩余:%6d 总量:%6d 进度:%3.2f%s"%(current,size,percent,"%")
            sys.stdout.write(msg)
            if(current==0):
                return None
        return r.get()