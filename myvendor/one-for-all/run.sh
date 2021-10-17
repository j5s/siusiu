#!/bin/bash
base_path=$HOME/src
install_path=$base_path/one-for-all

function download {
    git clone https://github.com.cnpmjs.org/shmilylty/OneForAll.git $1
    cd $1
    cd OneForAll/
    python3 -m pip install -U pip setuptools wheel -i https://mirrors.aliyun.com/pypi/simple/
    pip3 install -r requirements.txt -i https://mirrors.aliyun.com/pypi/simple/

}

#1.检查程序目录是否存在
if [ ! -d $install_path ]; then
    #2.如果不存在就下载
    echo "[*] download..."
    download $install_path
fi
python3 $install_path/oneforall.py $*

