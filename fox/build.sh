#!/bin/bash

# mac build
go build -o build/fox-mac main.go

# linux build
GOOS=linux GOARCH=amd64 go build -o build/fox main.go

# copy config.yaml
cp -f conf.yaml build/conf.yaml