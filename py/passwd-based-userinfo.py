# 基于用户名、邮箱、手机号生成弱口令
# 1. 总结常见账号
# 2. 总结常见密码
# 3. 密码：以常见密码为母本，结合密码组合规则生成
# 4. 扩展到BurpSuite

import exrex
base = ['123','1234','12345','888','666','88888888']

def generateRules():
    keys = ["(|{email})","(|{phone})","(|{username})","(|{base_word})"]
    rules = []
    for key1 in keys:
        for key2 in keys:
            for key3 in keys:
                for key4 in keys:
                    if(key1!=key2 and key1!=key3 and key1!=key4 and key2!=key3 and key2!=key4 and key3!=key4):
                        rules.append(key1+key2+key3+key4) 
    return rules

def main(username="",email="",phone=""):
    rules = generateRules()
    dic = []
    for rule in rules:
        for base_word in base:
            r = rule.format(username=username,email=email,phone=phone,base_word=base_word)
            dic.extend(list(exrex.generate(r)))
    for d in dic:
        if(len(d)>=8):
            print(d)

if __name__ == '__main__':
    username=input("# username:")
    email=input("# email:")
    phone=input("# phone:")
    main(username,email,phone)