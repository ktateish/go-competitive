#!/bin/sh -e


d=.build
rm -rf $d
mkdir -p $d
cp -p *.go2 $d
cd $d
rm -f *_test.go2 *_test.go

#export GO111MODULE=off
go2go build
# XXX: fix gottani's bug
sed -i -e 's|^//line .*:[0-9][0-9]*$||' *.go
gottani . > ../gottani.go
cd ..
go build -o a.out gottani.go debug.go
