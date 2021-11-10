#!/bin/bash
base_path=$HOME/src
install_path=$base_path/wafw00f #程序安装目录

function download {
    git clone https://github.com.cnpmjs.org/EnableSecurity/wafw00f.git $1
    cd $install_path && sudo python setup.py install 
}
#1.检查程序目录是否存在
if [ ! -d $install_path ]; then
    #2.如果不存在就下载
    echo "[*] download..."
    download $install_path
    echo "[*] download success"
fi

cd $install_path && sudo python setup.py install > /dev/null && wafw00f $*

