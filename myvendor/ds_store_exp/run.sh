#!/bin/bash
base_path=$HOME/src
install_path=$base_path/ds_store_exp #程序安装目录

function download {
    git clone https://github.com.cnpmjs.org/lijiejie/ds_store_exp.git $1
    cd $1
    pip install -r requirements.txt --target=$1
}

#1.检查程序目录是否存在
if [ ! -d $install_path ]; then
    #如果不存在就下载
    echo "[*] download..."
    download $install_path
    echo "[*] download success!"
fi
python $install_path/ds_store_exp.py $*
