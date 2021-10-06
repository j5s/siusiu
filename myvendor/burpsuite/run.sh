#!/bin/bash
base_path=$HOME/src
install_path=$base_path/burpsuite #程序安装目录

function download {
    wget -P $1 https://portswigger.net/burp/releases/download?product=pro &
    version=2020.11.3 &
    type=macosx
    git clone https://github.com.cnpmjs.org/TrojanAZhen/BurpSuitePro-2.1.git $1
}

#1.检查程序目录是否存在
if [ ! -d $install_path ]; then
    #2.如果不存在就下载
    echo "[*] download..."
    download $install_path
fi
echo "已经下载"
