#!/bin/sh -e

d=.fastbuild
mkdir -p $d
cat <<EOH > $d/main.go2
package main
// vim:set ft=go:

import (
	"bufio"
	"bytes"
	"encoding/binary"
	"fmt"
	"io"
	"math"
	"math/bits"
	"math/cmplx"
	"math/rand"
	"os"
	"runtime"
	"strconv"
	"strings"
	"unsafe"
)

EOH

MARK='^// snip ------------------------------------------------------------------------$'
grep -h -A 100000000 "$MARK" $(ls -1 *.go2 | grep -v '_test.go2$' | grep -v '^main.go2') | grep -Ev '^(// snip --|--$)' >> $d/main.go2
grep -h -A 100000000 "$MARK" main.go2 >> $d/main.go2
goimports -w $d/main.go2
cd $d

#export GO111MODULE=off
go2go build
# XXX: fix gottani's bug
sed -i -e 's|^//line .*:[0-9][0-9]*$||' *.go
gottani . > ../gottani.go
sed -i -e 's|^//line .*:[0-9][0-9]*$||' ../gottani.go
cd ..
go build -o a.out gottani.go debug.go
