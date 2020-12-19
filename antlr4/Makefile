GOFMT ?= gofmt -s -w
GOFILES := $(shell find . -name "*.go" -type f)

TARGET = antlr4
RELEASE_ROOT = release
RELEASE_LINUX_AMD64 = $(RELEASE_ROOT)/linux-amd64/$(TARGET)
RELEASE_DARWIN_AMD64 = $(RELEASE_ROOT)/darwin-amd64/$(TARGET)
RELEASE_DARWIN_ARM64 = $(RELEASE_ROOT)/darwin-arm64/$(TARGET)

.PHONY: build
build: fmt
	go build -o $(TARGET) main.go

.PHONY: release
release: linux-amd64 darwin-amd64 darwin-arm64
	cp ../LICENSE README.md $(RELEASE_LINUX_AMD64)
	cp ../LICENSE README.md $(RELEASE_DARWIN_AMD64)
	cp ../LICENSE README.md $(RELEASE_DARWIN_ARM64)
	cd $(RELEASE_LINUX_AMD64)/.. && rm -f *.zip && zip -r $(TARGET)-linux_amd64.zip $(TARGET) && cd -
	cd $(RELEASE_DARWIN_AMD64)/.. && rm -f *.zip && zip -r $(TARGET)-darwin_amd64.zip $(TARGET) && cd -
	cd $(RELEASE_DARWIN_ARM64)/.. && rm -f *.zip && zip -r $(TARGET)-darwin_arm64.zip $(TARGET) && cd -

.PHONY: linux-amd64
linux-amd64:
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o $(RELEASE_LINUX_AMD64)/$(TARGET) main.go

.PHONY: darwin-amd64
darwin-amd64:
	CGO_ENABLED=0 GOOS=darwin GOARCH=amd64 go build -o $(RELEASE_DARWIN_AMD64)/$(TARGET) main.go

.PHONY: darwin-arm64
darwin-arm64:
	CGO_ENABLED=0 GOOS=darwin GOARCH=arm64 go build -o $(RELEASE_DARWIN_ARM64)/$(TARGET) main.go

.PHONY: fmt
fmt:
	$(GOFMT) -w $(GOFILES)
