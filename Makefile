PROJECT_NAME := GoDDNSClient
BUILD_DIR := bin

.PHONY: all clean windows linux

all: windows linux

clean:
	rm -rf $(BUILD_DIR)

windows:
	GOOS=windows GOARCH=amd64 go build -o $(BUILD_DIR)/$(PROJECT_NAME).exe

linux: linux_amd64 linux_arm64

linux_amd64:
	GOOS=linux GOARCH=amd64 go build -o $(BUILD_DIR)/$(PROJECT_NAME)_linux_amd64

linux_arm64:
	GOOS=linux GOARCH=arm64 go build -o $(BUILD_DIR)/$(PROJECT_NAME)_linux_arm64 