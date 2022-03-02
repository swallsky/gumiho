#!/bin/bash

# mac build
go build -o build/tail-mac main.go

# linux build
GOOS=linux GOARCH=amd64 go build -o build/tail main.go

# copy config.yaml
cp -f conf.yaml build/conf.yaml

# 拷贝相应的命令
cp -rf script/ build/script/

# 更改权限
chmod +x script/*.sh