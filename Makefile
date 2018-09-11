all: install

ensure:
	dep ensure

install:
	go install -v ./...

build:
	env GOOS=linux CGO_ENABLED=1 go build -o builds/lappy cmd/lappy/main.go

test:
	go test -race -cover ./...

image:
	docker build --rm -t bitbrewers/lappy .

dev:
	docker build --rm -f Dockerfile.dev -t bitbrewers/lappy-dev .
	docker run \
		-e LOG_LEVEL='debug' \
		-e DSN='file:test.db' \
		-e LISTEN_ADDR=':8000' \
		-e SERIAL_PORT='ttyS11' \
		-v $(shell pwd):/go/src/github.com/bitbrewers/lappy \
		--rm -it bitbrewers/lappy-dev:latest

run:
	docker run  --rm -it bitbrewers/lappy

.PHONY: install test dev run
