BINARY_NAME=imotif-tools

all: build-linux build-mac build-windows

build-linux:
	GOOS=linux GOARCH=amd64 go build -o bin/$(BINARY_NAME)-linux cmd/imotif-tools/main.go

build-mac:
	GOOS=darwin GOARCH=amd64 go build -o bin/$(BINARY_NAME)-macos cmd/imotif-tools/main.go

build-windows:
	GOOS=windows GOARCH=amd64 go build -o bin/$(BINARY_NAME).exe cmd/imotif-tools/main.go
