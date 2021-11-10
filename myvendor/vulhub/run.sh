#!/bin/bash
base_path=$HOME/src
install_path=$base_path/vulhub #程序安装目录

function download {
    git clone https://github.com.cnpmjs.org/vulhub/vulhub.git $1
}

function show_help {

    cd $install_path && ls * 
    echo 
    echo "[*] vulhub所在目录:"
    pwd
    echo 
    echo "[*] Example:"
    echo 
    echo "vulhub 查看所有靶场"
    echo "vulhub thinkphp/2-rce 生成thinkphp/2-rce镜像并运行"
    echo && echo "[*] 在线文档:" && echo
    echo "https://vulhub.org/#/environments/"
}

#1.检查程序目录是否存在
if [ ! -d $install_path ]; then
    #如果不存在就下载
    echo "[*] download..."
    download $install_path
    echo "[*] download success!"
fi

case $# in
1)
    if [ ! -d $install_path/$1 ] || [ -z $1 ]; then
        show_help
        exit 0
    fi
    cd $install_path/$1
    echo "正在生成镜像中..." && docker-compose up -d && echo && echo "当前所有镜像:" && docker images && echo && echo "当前正在运行的容器:" && docker ps
    ;;
*)
    show_help
    ;;
esac
