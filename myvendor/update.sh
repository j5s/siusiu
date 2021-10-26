#!/bin/bash
path="$HOME/src/siusiu"
cd $path && git reset --hard && git pull origin master && go build ./... && go build && echo "更新已完成，请exit退出，重新进入siusiu"