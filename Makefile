BINARY_NAME=itgc

all: build-linux build-mac build-windows

build-linux:
	GOOS=linux GOARCH=amd64 go build -o bin/$(BINARY_NAME)-linux main.go

build-mac:
	GOOS=darwin GOARCH=amd64 go build -o bin/$(BINARY_NAME)-macos main.go

build-windows:
	GOOS=windows GOARCH=amd64 go build -o bin/$(BINARY_NAME).exe main.go
