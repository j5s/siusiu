#!/bin/bash
if [ $# -lt 1 -o ${#1} -eq 0 ]; then
    echo "Usage:run.sh 192.168.1.0/24"
    exit 1
else
   nmap -sn -PE --min-hostgroup 1024 --min-parallelism 1024 -oX nmap.xml $1
   python3 ./py/c-segment-scan/scanner.py 
fi 
