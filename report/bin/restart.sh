#!/bin/bash
server=smsgate-report

if [ "$1" = "" ]; then
    sudo mv config.product.json config.json
fi

if [ "$1" = "product" ]; then
    sudo mv config.product.json config.json
fi

if [ "$1" = "producttest" ]; then
    sudo mv config.producttest.json config.json
fi

if [ "$1" = "dev" ]; then
    sudo mv config.dev.json config.json
fi

sudo killall $server

sudo nohup ./$server config.json &

ps -ef | grep $server
