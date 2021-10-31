GO2SRCS = $(wildcard *.go2)

all: fastbuild

build: b.out

fastbuild: a.out

a.out: $(GO2SRCS)
	./.fastbuild.sh

b.out: $(GO2SRCS)
	./.build.sh

test:
	./.test.sh

test-v:
	./.test.sh -v

check: a.out
	./.check.sh

clean:
	rm -rf .build .fastbuild a.out gottani.go *.go-e

.PHONY: all build clean test update-template
