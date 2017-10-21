#!/bin/zsh

GOOS=linux GOARCH=amd64 go build -o duck -v main.go 2> /dev/null
