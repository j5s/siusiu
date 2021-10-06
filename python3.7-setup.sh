#!/bin/bash
setup_path=$HOME/src
if [ ! -d $setup_path ];then
    mkdir $setup_path
fi
wget -P $setup_path https://www.python.org/ftp/python/3.7.1/Python-3.7.1.tgz
tar -zxvf $setup_path/Python-3.7.1.tgz
cd $setup_path/Python-3.7.1
./configure
make
make install
