#!/bin/bash
. $HOME/src/siusiu/myvendor/lib/common.lib
install_path=$base_path/observer-ward
download_url="https://github.com.cnpmjs.org/0x727/ObserverWard_0x727/releases/download/default"

case $(get_os) in
"linux_amd64")
    exe="observer_ward_amd64"
    ;;
"mac")
    exe="observer_ward_darwin"
    ;;
*)
    echo "未知的操作系统类型"
    exit 1
    ;;
esac

download_url=$download_url/$exe
if [ ! -d $install_path ]; then
    git clone https://github.com.cnpmjs.org/0x727/FingerprintHub.git $install_path
    
    if [ ! -f $install_path/$exe ]; then
        wget -P $install_path $download_url
    fi
fi

if [ ! -x $install_path/$exe ]; then
    chmod +x $install_path/$exe
fi
if [ ! -f $HOME/bin/observer_ward ];then 
    cp $install_path/$exe $HOME/bin/observer-ward
fi
if [ ! -d $HOME/bin/fingerprint ];then
    cp -r $install_path/fingerprint $HOME/bin/fingerprint 
fi

$install_path/$exe $*
