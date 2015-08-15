SRC=./src/tests
GOPATH=$(shell pwd)

all: generate build test 

build:
	go build -v ${SRC}

generate:
	go generate -v -x ${SRC}

test:
	go test -v ${SRC}
