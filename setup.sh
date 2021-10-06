#!/bin/bash

setup_path="$HOME/src"
bin_path=$HOME/bin
app_name=siusiu
function setup {
    git clone https://gitee.com/nothing-is-nothing/siusiu.git $1/$app_name
    cd $1/$app_name
    go build -o $app_name
    mv $app_name $2/
    #删除含有app_name的所有行
    sed -i "/$app_name/d" $HOME/.zshrc
    sed -i "/$app_name/d" $HOME/.bash_profile
    #向.zshrc和.bash_profile 中添加别名
    echo "alias $app_name=$2/$app_name" >> $HOME/.zshrc
    echo "alias $app_name=$2/$app_name" >> $HOME/.bash_profile
    source $HOME/.zshrc
    source $HOME/.bash_profile
    echo "[*] setup success!"
}

if [ ! -d $setup_path ]; then
    mkdir $setup_path
fi

setup $setup_path $bin_path
