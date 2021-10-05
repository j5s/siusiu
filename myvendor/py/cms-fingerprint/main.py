from scancms import ScanCms
from optparse import OptionParser

def main(url,thread_count):
    scan_cms = ScanCms(url,thread_count)
    scan_cms.run()


if __name__ == "__main__":
    parser = OptionParser(description="CMS FingerPrint Scanner V1.0 Author:无在无不在 Email:2227627947@qq.com")
    parser.add_option("-u","--url",dest="url",help="target url eg:http://xdu.databankes.cn")
    parser.add_option("-t","--thread_count",dest="thread_count",help="thread count default 10",type='int',default=10)
    (options,args) = parser.parse_args()
    if(options.url):
        main(options.url,options.thread_count)
    else:
        parser.print_help()
