#!/bin/bash

function build(){
    CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o store-api-user
}

function run(){
    ./activity-new.exe
}


if [ "$1" == "run" ];then
    run
elif [ "$1" == "build" ];then
    build
elif [ "$1" == "build_l" ];then
    build_l
fi