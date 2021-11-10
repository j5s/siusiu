#!/bin/bash

setup_path="$HOME/src"
app_name=siusiu
shell_config_file_list="$HOME/.zshrc $HOME/.bash_profile $HOME/.bashrc /etc/profile"

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
        if [ $os = 'centos' -o $os = 'ubuntu' -o $os = 'linux' ]; then
            go_pkg="go1.15.5.linux-amd64.tar.gz"
            download_url="https://studygolang.com/dl/golang/$go_pkg"
            echo "检测当前目录是否有go语言安装包"
            if [ ! -f $HOME/$go_pkg ]; then
                echo "没有go语言安装包，开始下载"
                wget -P $HOME $download_url && rm -rf /usr/local/go && tar -C /usr/local -xzf $HOME/$go_pkg
            else
                echo "有go语言安装包，开始解压"
                rm -rf /usr/local/go && tar -C /usr/local -xzf $HOME/$go_pkg
            fi
        elif [ $os = 'mac' ]; then
            echo "检测到操作系统为mac,开始检测家目录下是否有安装包"
            go_pkg="go1.15.5.darwin-amd64.pkg"
            if [ ! -f $HOME/$go_pkg ]; then
                echo "家目录下没有安装包,开始下载"
                download_url="https://studygolang.com/dl/golang/$go_pkg"
                wget -P $HOME $download_url
            fi
            open $go_pkg
            while true; do
                if [ $(ps aux | grep "Installer.app" | grep -v grep | wc -l) -eq 0 ]; then
                    break
                else
                    sleep 1
                fi
            done
            echo "go语言安装包安装成功"
        fi

        for shell_config_file in $shell_config_file_list; do
            if [ -f $shell_config_file ]; then
                sed -i '/\/usr\/local\/go/d' $shell_config_file
                echo 'export PATH=$PATH:/usr/local/go/bin' >>$shell_config_file && source $shell_config_file && echo "重新加载 $shell_config_file 成功"
            fi
        done
        if [ $(go version | grep "1.15.5" | wc -l) -eq 1 ]; then
            echo "go安装成功"
        fi
    else
        echo "已安装go"
    fi
}
# 初始化go配置
function go_init {
    echo "正在初始化go配置..."
    #1.打开go mod 进行依赖管理
    go env -w GO111MODULE=on
    #2.设置第三方库的镜像代理
    go env -w GOPROXY=https://goproxy.cn,https://goproxy.io,direct
    #3. 下载goimports 工具
    go get -v golang.org/x/tools/cmd/goimports
    echo "go配置初始化完成"
}

function setup {
    #检查并安装go语言环境
    download_go && go_init
    git config --global http.sslverify false
    if [ -d $1/$app_name ]; then
        echo "已安装 siusiu,正在检查更新..."
        cd $1/$app_name && git reset --hard && git pull origin master && go build ./... && go build
    else
        echo "未安装 siusiu,正在下载中..."
        git clone --depth 1 https://github.com.cnpmjs.org/ShangRui-hash/siusiu.git $1/$app_name && cd $1/$app_name && go build ./... && go build -o $app_name
    fi

    for shell_config_file in $shell_config_file_list; do
        if [ -f $shell_config_file ]; then
            sed -i "/$app_name/d" $shell_config_file && echo "alias $app_name=$1/$app_name/$app_name" >> $shell_config_file && source $shell_config_file && echo "重新加载 $shell_config_file 成功"
        fi
    done
    echo "[*] setup success!"
}

if [ ! -d $setup_path ]; then
    mkdir $setup_path
fi
if [ ! -d $HOME/bin ];then 
    mkdir $HOME/bin
fi

setup $setup_path
