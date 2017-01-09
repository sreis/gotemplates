SRC=./tests
GOPATH=$(shell pwd)

all: generate format test 

format:
	gofmt -w ${SRC}

generate:
	go generate -v -x ${SRC}

test:
	go test -race -v ${SRC}
