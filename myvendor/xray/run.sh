#!/bin/bash
base_path=$HOME/src
install_path=$base_path/xray #程序安装目录

function download {
    if [ ! -f $1/$name.zip ]; then
        wget -P $1 "https://download.xray.cool/xray/1.7.1/$name.zip"
    fi
    cd $1
    unzip "$name.zip"
}

function get_os {
    case $(uname -s) in
    "Darwin")
        echo "mac"
        ;;
    "Linux")
        hardware_platform=$(uname -i)
        if [ hardware_platform='x86_64' ]; then
            echo "linux_amd64"
        elif [ hardware_platform="i386" ]; then
            echo "linux_386"
        fi
        ;;
    *)
        echo "unkown"
        ;;
    esac

}

case $(get_os) in
"mac")
    name="xray_darwin_amd64"
    ;;
"linux_386")
    name="xray_linux_386"
    ;;
"linux_amd64")
    name="xray_linux_amd64"
    ;;
*)
    echo "未知的操作系统类型"
    exit 1
    ;;
esac
#1.检查程序目录是否存在
if [ ! -d $install_path  ] || [ ! -f  $install_path/$name ]; then
    #如果不存在就下载
    echo "[*] download..."
    download $install_path
    echo "[*] download success!"
fi
cd $install_path && echo "当前目录:$(pwd)"
$install_path/$name $*
