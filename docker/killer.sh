#!/bin/bash
Y1="\e[1;33m"; # Yellow
RS="\e[0;00m" # Reset
echo "${Y1} ==> Execute Killer <== ${RS}"
PID=`ps aux | grep service | grep -v "grep" | awk '{print $2}'`
kill -9 $PID
cd /var/app ; make build
/var/app/bin/service &
