#!/bin/sh -e


d=.build
rm -rf $d
mkdir -p $d
cp -p *.go2 $d
cd $d

#export GO111MODULE=off
go2go build
sed -i -e 's|^//line .*:[0-9][0-9]*$||' *.go
go test "$@"
