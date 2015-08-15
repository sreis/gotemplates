SRC=./...
GOPATH=$(shell pwd)

build:
	go build -v ${SRC}

generate:
	go generate -v -x ${SRC}

test:
	go test -v ${SRC}
