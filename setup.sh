#!/bin/bash

setup_path="$HOME/src"
app_name=siusiu

function get_os {
    a=$(uname -a)
    b="Darwin"
    c="centos"
    d="ubuntu"
    e="Linux"
    if [[ $a =~ $b ]]; then
        echo "mac"
    elif [[ $a =~ $c ]]; then
        echo "centos"
    elif [[ $a =~ $d ]]; then
        echo "ubuntu"
    elif [[ $a =~ $e ]]; then
        echo "linux"
    else
        echo $a
    fi
}

function download_go {
    echo "正在检查是否安装了go语言环境"
    if [ $(go version | wc -l) -ne 1 ]; then
        echo "未安装go"
        os=$(get_os)
        if [ ($os -eq 'centos' || $os -eq 'ubuntu' || $os -eq 'linux' ]; then
            go_pkg="go1.15.5.linux-amd64.tar.gz"
            download_url="https://studygolang.com/dl/golang/$go_pkg"
            wget $download_url && rm -rf /usr/local/go && tar -C /usr/local -xzf $go_pkg && rm $go_pkg
        elif [ $os -eq 'mac' ]; then
            go_pkg="go1.15.5.darwin-amd64.pkg"
            download_url="https://studygolang.com/dl/golang/$go_pkg"
            wget $download_url && open $go_pkg
        fi

        for shell_config_file in "$HOME/.zshrc" "$HOME/.bash_profile" "$HOME/.bashrc"; do
            if [ -f $shell_config_file ]; then
                sed -i '/\/usr\/local\/go/d' $shell_config_file
                echo 'export PATH=$PATH:/usr/local/go/bin' >>$shell_config_file #向shell配置文件中添加别名
                source $shell_config_file
            fi
        done
        if [ $(go version | grep "1.15.5" | wc -l) -eq 1 ]; then
            echo "go安装成功"
        fi
    else
        echo "已安装go"
    fi
}

function setup {
    git config --global http.sslverify false
    rm -rf $1/$app_name && git clone https://gitee.com/nothing-is-nothing/siusiu.git $1/$app_name
    cd $1/$app_name
    download_go
    go build -o $app_name
    for shell_config_file in "$HOME/.zshrc" "$HOME/.bash_profile" "$HOME/.bashrc"; do
        if [ -f $shell_config_file ]; then
            sed -i "/$app_name/d" $shell_config_file                           #删除含有app_name的所有行
            echo "alias $app_name=$1/$app_name/$app_name" >>$shell_config_file #向shell配置文件中添加别名
            source $shell_config_file
        fi
    done
    echo "[*] setup success!"
}



if [ ! -d $setup_path ]; then
    mkdir $setup_path
fi

setup $setup_path
