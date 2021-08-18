.PHONY: get clean build-linux build-windows build
.SILENT: ${.PHONY}

APP_NAME=traffic-generator

all: build

get:
	go get ./...

clean:
	rm -rf *.log bin
	go clean -cache

build-linux: get
	GOOS=linux GOARCH=386 go build -o bin/${APP_NAME}-linux-386
	GOOS=linux GOARCH=amd64 go build -o bin/${APP_NAME}-linux-amd64

build-windows: get
	GOOS=windows GOARCH=386 go build -o bin/${APP_NAME}-windows-386.exe
	GOOS=windows GOARCH=amd64 go build -o bin/${APP_NAME}-windows-amd64.exe

build: get clean build-linux build-windows
