.PHONY: build install check clean vet

all: build install

build:
	go build .

install:
	go install .

check: build
	./lazybeads --check

vet:
	go vet ./...

clean:
	go clean
	rm -f lazybeads
