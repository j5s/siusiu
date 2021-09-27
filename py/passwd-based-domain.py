'''
Date: 2021-01-17 21:58:28
LastEditors: 无在无不在
LastEditTime: 2021-01-18 20:28:25
'''
import exrex 
import sys
"""
基于域名生成字典：
域名可能的形式：
demo.webdic.com
http://demo.webdic.com
http://demo.webdic.com/
https://demo.webdic.com
https://demo.webdic.com/
生成思路：
demo webdic 这类的关键字都可能成为密码的一部分
"""
useless = ['com','cn','org','edu','www','gov']
base = ['admin','guest','root','manage','manager','123','1234','12345','888','666','88888888']

def hostParser(host):
    """去除协议名"""
    if '://' in host:
        host = host.split("://")[1].replace('/', '')
    else:
        host = host.replace("/", '')
    return host
def dictGenerator(host,rule):
    dict = [ ]
    key_words = host.split('.')
    for k in key_words:
        if k not in useless:
            for b in base:
                dict.extend(list(exrex.generate(rule.format(key_word=k,base_word=b))))
    
    return dict

def main():
    domain= input("pls input doamin:")
    host = hostParser(domain)
    rule = "(|{key_word})(|!|@|#)(|{base_word})(|!|@|#)(|20[0-2][0-9])"

    dict = dictGenerator(host,rule)
    file = open("passwd_1.txt", "a+")
    for d in dict:
        if(len(d)>=8):
            print(d)
            file.write(d+'\n')
    file.close()
    

if __name__ == "__main__":
    main()


