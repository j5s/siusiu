#!/bin/bash
base_path=$HOME/src
install_path=$base_path/Glass #程序安装目录

function download {
    git clone https://github.com.cnpmjs.org/s7ckTeam/Glass $1
    cd $1
    pip3 install -r requirements.txt
}

#1.检查程序目录是否存在
if [ ! -d $install_path ]; then
    #如果不存在就下载
    echo "[*] download..."
    download $install_path
    echo "[*] download success!"
fi
python3 $install_path/Glass.py $*
