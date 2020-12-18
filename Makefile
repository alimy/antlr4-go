GOFMT ?= gofmt -s -w
GOFILES := $(shell find . -name "*.go" -type f)

.PHONY: default
default: ci

.PHONY: ci
ci:
	go test ./...

.PHONY: test
test: fmt
	go test ./...

.PHONY: vet
vet:
	go vet ./...

.PHONY: fmt
fmt:
	$(GOFMT) $(GOFILES)
