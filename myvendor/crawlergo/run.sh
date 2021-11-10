#!/bin/bash
base_path=$HOME/src
install_path=$base_path/crawlergo #crawlergo安装目录
chromium_path=$base_path/chromium #chromium 安装目录
chromium_mac_download_link=https://npm.taobao.org/mirrors/chromium-browser-snapshots/Mac/901912/chrome-mac.zip
chromium_linux_download_link=https://npm.taobao.org/mirrors/chromium-browser-snapshots/Linux_x64/901912/chrome-linux.zip

function get_os {
    a=$(uname -a)
    if [[ $a =~ Darwin ]]; then
        echo "mac"
    elif [[ $a =~ Linux ]] || [[ $a =~ ubuntu ]] || [[ $a =~ centos ]]; then
        echo "linux"
    else
        echo $a
    fi
}

function download_crawlergo {
    if [ ! -d $install_path ]; then
        echo "[*] download..."
        git clone https://github.com.cnpmjs.org/Qianlitp/crawlergo.git $install_path && cd $install_path/cmd/crawlergo && go build ./... && go build
        echo "[*] download success"
    fi
}

function download_chromium {
    case $(get_os) in
    mac)
        if [ ! -d $chromium_path ]; then
            echo "[*] download chromium ..."
            wget -P $chromium_path $chromium_mac_download_link
            echo "[*] download chromium success"
        fi
        if [ ! -e $chromium_path/chrome ]; then
            unzip $chromium_path/chrome-mac.zip -d $chromium_path/chrome
        fi
        ;;
    linux)
        if [ ! -d $chromium_path ]; then
            wget -P $chromium_path $chromium_linux_download_link
        fi
        ;;
    *)
        echo "暂时不支持该操作系统"
        ;;
    esac

}

#1.检查程序目录是否存在
current=$(pwd)

#2.如果不存在就下载
download_chromium
download_crawlergo
#3.检查更新
cd $install_path && git reset --hard && git pull origin master && cd ./cmd/crawlergo && go build ./... && go build && cd $current

#4.运行
case $(get_os) in
mac)
    $install_path/cmd/crawlergo/crawlergo -c $chromium_path/chrome/chrome-mac/Chromium.app/Contents/MacOS/Chromium --custom-headers "{\"User-Agent\": \"Mozilla/5.0 (Windows NT 6.1; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/79.0.3945.0 Safari/537.36\"}" $*
    ;;
linux)
    $install_path/cmd/crawlergo/crawlergo $*
    ;;
*)
    echo "暂时不支持该操作系统"
    ;;
esac
