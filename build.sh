#!/bin/sh

IFS='/'
go tool dist list | while read os arch; do
    echo "env GOOS=$os GOARCH=$arch\tgo build -v -o logstamper.${os}_${arch}\tgithub.com/tokenshift/logstamper"
    env GOOS=$os GOARCH=$arch go build -o logstamper.${os}_${arch} github.com/tokenshift/logstamper
done