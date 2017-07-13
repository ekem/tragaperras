BINARY = juego

ifeq ($(OS),Windows_NT)
	BINARY := $(BINARY).exe
endif

all:
	go build -o bin/$(BINARY) src/juego.go

clean:
	@(if [ -d bin/ ] ; then rm -r bin/* bin/ ; fi)
	@(if [ -f nativelog.txt ] ; then rm nativelog.txt ; fi)

PHONY: clean all
