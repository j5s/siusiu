#!/bin/bash
base_path=$HOME/src
dirsearch_path=$base_path/dirsearch

function download {
    git clone https://github.com.cnpmjs.org/maurosoria/dirsearch.git $1
    cd $1
    pip3 install -r requirements.txt
}

#1.检查程序目录是否存在
if [ ! -d $dirsearch_path ]; then
    #2.如果不存在就下载
    echo "[*] download dirsearch..."
    download $dirsearch_path
fi
python3 $dirsearch_path/dirsearch.py $*

