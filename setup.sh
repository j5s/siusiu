#!/bin/bash

setup_path="$HOME/src"
bin_path=$HOME/bin
app_name=siusiu
function setup {
    git clone https://gitee.com/nothing-is-nothing/siusiu.git $1/$app_name
    cd $1/$app_name
    go build -o $app_name
    mv $app_name $2/
    echo "alias $app_name=$2/$app_name" >> $HOME/.zshrc
    source $HOME/.zshrc
    echo "[*] setup success!"
}

if [ ! -d $setup_path ]; then
    mkdir $setup_path
fi

setup $setup_path $bin_path
