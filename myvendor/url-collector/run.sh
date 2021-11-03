#!/bin/bash
base_path=$HOME/src
install_path=$base_path/url-collector #程序安装目录

function download {
    git clone https://gitee.com/nothing-is-nothing/url-collector.git $1 && cd $1 && go build ./... && go build
}

#1.检查程序目录是否存在
if [ ! -d $install_path ]; then
    #2.如果不存在就下载
    echo "[*] download..."
    download $install_path
    echo "[*] download success"
fi
current=$(pwd)
cd $install_path && git reset --hard && git pull origin master && go build ./... && go build && cd $current

if [ ! -x $install_path/url-collector ];then
    chmod +x $install_path/url-collector
fi

if [ ! -f $HOME/bin/url-collector ];then
    cp $install_path/url-collector $HOME/bin/ 
fi

$install_path/url-collector $*
