TARGETS="windows/amd64,windows/386,darwin/amd64,darwin/386,linux/amd64,linux/386"
LDFLAGS="-X main.VERSION=`git describe --tags` -s -linkmode external -extldflags -static -w"
PKG="github.com/elwinar/rambler"

default: build
all: fetch build test

.PHONY: build
build: ## Build the binary for the local architecture
	@go build --ldflags=$(LDFLAGS) -o dist/rambler

.PHONY: fetch
fetch: ## Fetch the dependencies
	go get -d ./...

.PHONY: help
help: ## Get help
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | sort | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-10s\033[0m %s\n", $$1, $$2}'

.PHONY: image
image: ## Build the docker image
	docker build . -t elwinar/rambler --no-cache

.PHONY: release
release: ## Build the release files
	xgo --targets=$(TARGETS) --ldflags=$(LDFLAGS) $(PKG)

.PHONY: test
test: ## Test the project
	go test ./...
