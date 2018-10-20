all: install

install:
	go install -v ./...

build:
	env GOOS=linux CGO_ENABLED=1 go build -o builds/lappy cmd/lappy/main.go

test:
	go test -race -cover ./...

image:
	docker build --rm -t bitbrewers/lappy .

dev-image:
	docker build -f etc/Dockerfile --rm -t bitbrewers/lappy:dev .

run-dev: dev-image
	docker run \
		-p 8000:8000 \
		-e LOG_LEVEL='debug' \
		-e DSN='file:/tmp/test.db' \
		-e LISTEN_ADDR=':8000' \
		-v $(shell pwd):/lappy \
		--rm -it bitbrewers/lappy:dev \

publish-images:
	docker push bitbrewers/lappy:latest
	docker push bitbrewers/lappy:dev

.PHONY: install test run run-dev
