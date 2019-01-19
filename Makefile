.DEFAULT_GOAL			:= help
APP_NAME 				= lappy
CONTAINER_NAME 			= bitbrewers/${APP_NAME}

export GOFLAGS = -mod=vendor
export GO111MODULE = on

help:  ## Show help
	@awk 'BEGIN {FS = ":.*##"; printf "\nUsage:\n  make \033[36m<target>\033[0m\n\nTargets:\n"} /^[a-zA-Z_-]+:.*?##/ { printf "  \033[36m%-16s\033[0m %s\n", $$1, $$2 }' $(MAKEFILE_LIST)

all: test install ## Test and install lappy

.PHONY: install
install: ## Install lappy
	go install -v ./...

.PHONY: build
build: ## Build binary files
	env CGO_ENABLED=1 go build -o builds/lappy cmd/lappy/main.go

.PHONY: test
test: ## Run all tests including integration tests
	go test -race -cover ./...

.PHONY: lint
lint: ## Run static analyzer for whole code base. https://staticcheck.io
	staticcheck ./...

.PHONY: dev-image
dev-image: ## Build dev image
	docker build -f etc/Dockerfile --rm -t bitbrewers/lappy:dev .

.PHONY: run-image
run-dev: dev-image
	docker run \
		-p 8000:8000 \
		-e LOG_LEVEL='debug' \
		-e DSN='file:/tmp/test.db' \
		-e LISTEN_ADDR=':8000' \
		-v $(shell pwd):/lappy \
		--rm -it bitbrewers/lappy:dev \

.PHONY: vendor
vendor: ## Update vendor folder to match dependecies
	go mod vendor
