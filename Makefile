BINARY = juego

ifeq ($(OS),Windows_NT)
	APPEND = .exe
endif

all: jugar juego

jugar:
	go build -o bin/$@$(APPEND) src/$@.go
juego:
	go build -o bin/$@$(APPEND) src/$@.go

clean:
	@(if [ -d bin/ ] ; then rm -r bin/* bin/ ; fi)
	@(if [ -f nativelog.txt ] ; then rm nativelog.txt ; fi)

PHONY: clean all
