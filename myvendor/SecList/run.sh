#!/bin/bash
base_path=$HOME/src
install_path=$base_path/SecList #程序安装目录

function download {
    wget -P $1 -c https://github.com.cnpmjs.org/danielmiessler/SecLists/archive/master.zip -O SecList.zip && unzip $1/SecList.zip && rm -f $1/SecList.zip
}

#1.检查程序目录是否存在
if [ ! -d $install_path ]; then
    #如果不存在就下载
    echo "[*] download..."
    download $install_path
    echo "[*] download success!"
fi
ls $install_path
