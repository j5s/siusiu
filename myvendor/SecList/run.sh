#!/bin/bash
base_path=$HOME/src
install_path=$base_path/SecList #程序安装目录

function download {
    wget -c https://github.com.cnpmjs.org/danielmiessler/SecLists/archive/master.zip -O $1/SecList.zip 
    cd $1
    unzip SecList.zip -d SecList
    rm -f SecList.zip
}

#1.检查程序目录是否存在
if [ ! -d $install_path ]; then
    #如果不存在就下载
    echo "[*] download..."
    download $base_path
    echo "[*] download success!"
fi
# ls -lR $install_path | more
vim $install_path