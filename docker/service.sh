#!/bin/bash
base_app=/var/app
make build
$base_app/bin/service &
valuePaths="application cmd domain infra interfaces"
valueInit=`cd $base_app ; find $valuePaths -type f \( -name "*.go" \) -exec md5sum '{}' \;`
while true
  do
    valueNow=`cd $base_app ; find $valuePaths -type f \( -name "*.go" \) -exec md5sum '{}' \;`
    if [ "$valueInit" != "$valueNow" ]
      then
        valueInit="$valueNow"
        sh $base_app/docker/killer.sh
    fi
    sleep 2
  done
