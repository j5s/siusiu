#!/bin/bash
base_path=$HOME/src
install_path=$base_path/Vulcan #程序安装目录

function download {
    git clone https://github.com.cnpmjs.org/RedTeamWing/Vulcan.git $1 
}
#1.检查程序目录是否存在
if [ ! -d $install_path ]; then
    #2.如果不存在就下载
    echo "[*] download..."
    download $install_path
    echo "[*] download success"
fi
cd $1 && docker build .  && docker run -itd -p 8080:5000 VulnScan 

