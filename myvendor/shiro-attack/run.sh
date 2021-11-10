#!/bin/bash
base_path=$HOME/src
install_path=$base_path/shiro-attack #程序安装目录

function download {
    wget -P $1 https://github.com.cnpmjs.org/j1anFen/shiro_attack/releases/download/2.2/shiro_attack_2.2.zip
    cd $1
    unzip shiro_attack_2.2.zip -d ./ 
}

#1.检查程序目录是否存在
if [ ! -d $install_path ]; then
    #2.如果不存在就下载
    echo "[*] download..."
    download $install_path
fi

java -jar $install_path/shiro_attack_2.2/shiro_attack-2.2.jar > /dev/null 2>&1 &
