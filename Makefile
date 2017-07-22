BINARY = juego

ifeq ($(OS),Windows_NT)
	APPEND = .exe
	GOPATH = $(shell cygpath -w $(shell pwd))
endif

all: juego
	@(export GOPATH="$(GOPATH)")

jugar:
	go build -o bin/$@$(APPEND) src/$@.go
juego:
	go build -o bin/$@$(APPEND) src/$@.go

clean:
	@(if [ -d bin/ ] ; then rm -r bin/* bin/ ; fi)

PHONY: clean all
