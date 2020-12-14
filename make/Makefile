include common/go.mk

.PHONY: all
all: ## build pipeline
all: clean fmt test

.PHONY: clean
clean: ## remove files created during build
	$(call print-target)
	rm -fv coverage.*

.PHONY: test
test: ## go test with race detector and code covarage
	$(call print-target)
	EXIT=0 ;\
	go test -race -covermode=atomic -coverprofile=coverage.out ./... || EXIT=$$? ;\
	go tool cover -html=coverage.out -o coverage.html ;\
	exit $$EXIT
