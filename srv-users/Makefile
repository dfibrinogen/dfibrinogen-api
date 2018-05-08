PROJECT_NAME := "srv-events"
PKG := "github.com/dfibrinogen/dfibrinogen-api/$(PROJECT_NAME)"

PKG_LIST := $(shell go list ${PKG}/... | grep -v /vendor/)
GO_FILES := $(shell find . -name '*.go' | grep -v /vendor/ | grep -v _test.go)

FORMAT := go fmt $$(go list ./... | grep -v /vendor/)
CHECK := go vet $$(go list ./... | grep -v /vendor/)
TESTS := go test -race $$(go list ./... | grep -v /vendor/)

.PHONY: all dep build clean test coverage coverhtml lint

all: build

lint: ## Lint the files
	@golint -set_exit_status ${PKG_LIST}

test: ## Run unit tests
	${FORMAT}
	${CHECK}
	${TESTS}

race: dep ## Run data race detector
	@go test -race -short ${PKG_LIST}

msan: dep ## Run memory sanitizer
	@go test -msan -short ${PKG_LIST}

coverage: ## Generate global code coverage report
	.././coverage.sh;

coverhtml: ## Generate global code coverage report in HTML
	.././coverage.sh html;

dep: ## Get the dependencies
	@go get -v -d ./...
	@go get -u github.com/golang/lint/golint

build: dep ## Build the binary file
	@go build -i -v $(PKG)

clean: ## Remove previous build
	@rm -f $(PROJECT_NAME)

help: ## Display this help screen
	@grep -h -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'
