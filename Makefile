APP         = coincheck
VERSION     = $(shell git describe --tags --abbrev=0)
GIT_REVISION := $(shell git rev-parse HEAD)
GO          = go
GO_BUILD    = $(GO) build
GO_TEST     = $(GO) test
GO_TOOL     = $(GO) tool
GOOS        = ""
GOARCH      = ""
GO_PKGROOT  = ./...
GO_PACKAGES = $(shell $(GO_LIST) $(GO_PKGROOT))
GO_LDFLAGS  = 

.DEFAULT_GOAL := help
help: ## Show help message  
	@grep -E '^[0-9a-zA-Z_-]+[[:blank:]]*:.*?## .*$$' $(MAKEFILE_LIST) | sort \
	| awk 'BEGIN {FS = ":.*?## "}; {printf "\033[1;32m%-15s\033[0m %s\n", $$1, $$2}'

.PHONY: tools
tools: ## Install dependency tools 
	go install github.com/golangci/golangci-lint/cmd/golangci-lint@latest

.PHONY: test 
test: ## Run unit test
	env GOOS=$(GOOS) $(GO_TEST) -cover -coverpkg=$(GO_PKGROOT) -coverprofile=coverage.out $(GO_PKGROOT)
	$(GO_TOOL) cover -html=coverage.out -o coverage.html

.PHONY: clean
clean: ## Clean project
	-rm -rf coverage.out coverage.html

.PHONY: lint
lint: ## Run lint
	golangci-lint --config .golangci.yml run 
