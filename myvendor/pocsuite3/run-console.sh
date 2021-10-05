#!/bin/bash
base_path=$HOME/bin
install_path=$base_path/pocsuite3 #程序安装目录

. $base_path/biu/myvendor/pocsuite3/lib/download.lib

#1.检查程序目录是否存在
if [ -d $install_path ]; then
    python3 $install_path/pocsuite3-master/pocsuite3/console.py $*
else
    #2.如果不存在就下载
    echo "[*] download..."
    download $install_path
fi
