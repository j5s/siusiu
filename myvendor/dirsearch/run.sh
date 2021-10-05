#!/bin/bash
base_path=$HOME/bin
dirsearch_path=$base_path/dirsearch

function download
{
    git clone https://github.com.cnpmjs.org/maurosoria/dirsearch.git $1
    cd $1
    pip3 install -r requirements.txt
}


#1.检查程序目录是否存在
if [ -d $dirsearch_path ]
then
    python3 $dirsearch_path/dirsearch.py $*
else
#2.如果不存在就下载
    echo "[*] download dirsearch..."
    download $dirsearch_path
fi



   


# 设想：做过一个安全工具管理箱
# install.sh 自动下载
# run.sh 自动运行 -d 后台模式
# clear.sh 删除

# help 列出所有工具，用户可以选择需要的工具，比如：
# biu> dirsearch install 就可以自动下载
# biu> dirsearch 就可以自动运行
# biu> dirsearch clear 就可以自动删除
