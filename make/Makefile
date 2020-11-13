.PHONY: all
all: ## build pipeline
all: clean fmt test

.PHONY: clean
clean: ## remove files created during build
	$(call print-target)
	rm -fv coverage.*

.PHONY: fmt
fmt: ## go fmt
	$(call print-target)
	go fmt ./...

.PHONY: test
test: ## go test with race detector and code covarage
	$(call print-target)
	EXIT=0 ;\
	go test -race -covermode=atomic -coverprofile=coverage.out ./... || EXIT=$$? ;\
	go tool cover -html=coverage.out -o coverage.html ;\
	exit $$EXIT

### common stuff

.DEFAULT_GOAL := help

.PHONY: help
help:
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | sort | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'

define print-target
    @printf "Executing target: \033[36m$@\033[0m\n"
endef
