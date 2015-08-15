SRC=./tests
GOPATH=$(shell pwd)

all: generate test 

generate:
	go generate -v -x ${SRC}

test:
	go test -v ${SRC}
