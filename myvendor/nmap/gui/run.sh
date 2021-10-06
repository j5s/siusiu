#!/bin/bash
base_path=$HOME/src
install_path=$base_path/nmap      #安装程序目录
exe_path=/Applications/Zenmap.app #源码目录

function download {
    wget -P $install_path https://nmap.org/dist/nmap-7.92.dmg
    open $install_path/nmap-7.92.dmg
}

#1.检查程序目录是否存在
if [ !-d $exe_path ]; then
    echo "[*] download..."
    download $install_path
    echo "[*] download success!"
fi

open $exe_path
