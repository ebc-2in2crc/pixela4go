.DEFAULT_GOAL := help

GOCMD := env GO111MODULE=on go
GOMOD := $(GOCMD) mod
GOCLEAN := $(GOCMD) clean
GOINSTALL := $(GOCMD) install
GOTEST := $(GOCMD) test

.PHONY: deps
## Install dependencies
deps:
	$(GOMOD) download

.PHONY: devel-deps
## Install dependencies for develop
devel-deps: deps
	$(GOINSTALL) golang.org/x/tools/cmd/goimports@latest
	$(GOINSTALL) golang.org/x/lint/golint@latest
	$(GOINSTALL) github.com/Songmu/make2help/cmd/make2help@latest

.PHONY: test
## Run tests
test: deps
	$(GOTEST) -v ./...

.PHONY: lint
## Lint
lint: devel-deps
	go vet ./...
	golint -set_exit_status ./...

.PHONY: fmt
## Format source codes
fmt: devel-deps
	find . -name "*.go" -not -path "./vendor/*" | xargs goimports -w

.PHONY: clean
clean:
	$(GOCLEAN)
	rm -rf $(BUILDDIR)

.PHONY: help
## Show help
help:
	@make2help $(MAKEFILE_LIST)
