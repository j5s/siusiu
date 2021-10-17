#!/bin/bash
base_path=$HOME/src
install_path=$base_path/arp-spoofing #程序安装目录

function download {
    git clone https://gitee.com/nothing-is-nothing/arp-spoofing.git $1
    cd $1
    go build
}

#1.检查程序目录是否存在
if [ ! -d $install_path ]; then
    #2.如果不存在就下载
    echo "[*] download..."
    download $install_path
    echo "[*] download success"
fi

sudo $install_path/ARPSpoofing $*
