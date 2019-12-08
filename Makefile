.PHONY: all
all: build

MP_QUOTES_ENV ?= dev

clean:
	rm -rf bin/*

dependencies:
	go get -u github.com/codegangsta/negroni
	go get -u github.com/mattn/go-sqlite3
	go get -u github.com/gorilla/context
	go get -u github.com/gorilla/mux

build: dependencies build-api

build-api:
	go build -tags $(MP_QUOTES_ENV) -o ./bin/api api/main.go
