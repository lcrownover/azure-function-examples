.PHONY: all install clean run container handler

all:
	@go build -o bin/timer-function cmd/timer-function/main.go

run:
	@go run cmd/timer-function/main.go

install:
	@cp bin/timer-function /usr/local/bin/timer-function

container:
	@docker build -t timer-function .

handler:
	@go build -o handler cmd/timer-function/main.go

clean:
	@rm -f bin/timer-function /usr/local/bin/timer-function

