#!/bin/bash
base_path=$HOME/src
install_path=$base_path/nmap-cli #程序安装目录

function download {
    git clone https://github.com.cnpmjs.org/nmap/nmap.git $1
    cd $1
    ./configure
    make
    make install
}

#1.检查程序目录是否存在
if [ ! -d $install_path ]; then
    #2.如果不存在就下载
    echo "[*] download..."
    download $install_path
fi
sudo $install_path/nmap $*
