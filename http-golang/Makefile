.PHONY: all install clean run container handler

all:
	@go build -o bin/http-function cmd/http-function/main.go

run:
	@go run cmd/http-function/main.go

install:
	@cp bin/http-function /usr/local/bin/http-function

container:
	@docker build -t http-function .

handler:
	@go build -o handler cmd/http-function/main.go

clean:
	@rm -f bin/http-function /usr/local/bin/http-function
