#!/bin/bash

setup_path="$HOME/src"
bin_path=$HOME/bin
app_name=siusiu
function setup {
    git config --global http.sslverify false
    git clone https://gitee.com/nothing-is-nothing/siusiu.git $1/$app_name
    cd $1/$app_name
    go build -o $app_name
    mv $app_name $2/
    for shell_config_file in "$HOME/.zshrc" "$HOME/.bash_profile" "$HOME/.bashrc"; do
        if [ -f $shell_config_file ]; then
            sed -i "/$app_name/d" $shell_config_file                 #删除含有app_name的所有行
            echo "alias $app_name=$2/$app_name" >>$shell_config_file #向shell配置文件中添加别名
            source $shell_config_file
        fi
    done
    echo "[*] setup success!"
}

if [ ! -d $setup_path ]; then
    mkdir $setup_path
fi

setup $setup_path $bin_path
