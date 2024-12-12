VERSION=v0.1.0
GO_BUILD=go build -ldflags "-s -w -X github.com/TCP404/esdump/constant.VERSION=$(VERSION)" -o

BINARY_NAME=esdump
MAC_AMD_EXE :=./exe/$(BINARY_NAME)_amd64.dmg
MAC_ARM_EXE :=./exe/$(BINARY_NAME)_arm64.dmg
LINUX_EXE   :=./exe/$(BINARY_NAME)
WINDOWS_EXE :=./exe/$(BINARY_NAME).exe

BUILD_MAC_AMD64     := CGO_ENABLED=0 GOOS=darwin  GOARCH=amd64 $(GO_BUILD) $(MAC_AMD_EXE) main.go
BUILD_MAC_ARM64     := CGO_ENABLED=0 GOOS=darwin  GOARCH=arm64 $(GO_BUILD) $(MAC_ARM_EXE) main.go
BUILD_LINUX         := CGO_ENABLED=0 GOOS=linux   GOARCH=amd64 $(GO_BUILD) $(LINUX_EXE)   main.go
BUILD_WINDOWS       := CGO_ENABLED=0 GOOS=windows GOARCH=amd64 $(GO_BUILD) $(WINDOWS_EXE) main.go

COMPRESS_UPX=upx --best
COMPRESS_MAC_AMD := $(COMPRESS_UPX) $(MAC_AMD_EXE)
COMPRESS_LINUX   := $(COMPRESS_UPX) $(LINUX_EXE)
COMPRESS_WINDOWS := $(COMPRESS_UPX) $(WINDOWS_EXE)

start: clean build build-win build-linux build-mac compress compress-win compress-linux compress-mac

build:
	mkdir -p ./exe
	$(BUILD_MAC_ARM64)
	$(BUILD_MAC_AMD64)
	$(BUILD_LINUX)
	$(BUILD_WINDOWS)

build-win:
	mkdir -p ./exe
	$(BUILD_WINDOWS)

build-linux:
	mkdir -p ./exe
	$(BUILD_LINUX)

build-mac:
	mkdir -p ./exe
	$(BUILD_MAC_ARM64)
	$(BUILD_MAC_AMD64)

compress:
	$(COMPRESS_MAC_AMD)
	$(COMPRESS_LINUX)
	$(COMPRESS_WINDOWS)

compress-win:
	$(COMPRESS_WINDOWS)

compress-linux:
	$(COMPRESS_LINUX)
	
compress-mac:
	$(COMPRESS_MAC_AMD)


.PHONY: clean
clean:
	rm -rf ./exe/*

