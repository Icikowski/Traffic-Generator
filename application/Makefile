.PHONY: get clean build build-static build-all
.SILENT: ${.PHONY}

GIT_COMMIT := $(shell git rev-parse --short HEAD)
GIT_TAG := $(shell echo $${CURRENT_TAG:-$$(git describe --abbrev=0 | sed "s/v//")})

BASE_FLAGS := -X 'main.version=${GIT_TAG}' -X 'main.gitCommit=${GIT_COMMIT}'

all: build

get:
	go mod download -x

clean:
	rm -rf *.log bin
	go clean -cache

DYNAMIC_FLAGS := -ldflags "${BASE_FLAGS} -X 'main.binaryType=dynamic'"

build: clean
	go build ${DYNAMIC_FLAGS} .

STATIC_FLAGS := -ldflags "${BASE_FLAGS} -X 'main.binaryType=static' -w -extldflags '-static'"

build-static: clean
	env CGO_ENABLED=0 go build ${STATIC_FLAGS} .

build-all: clean
	env CGO_ENABLED=0 GOOS=windows GOARCH=amd64 go build ${STATIC_FLAGS} -o bin/traffic-generator-windows-amd64.exe .
	env CGO_ENABLED=0 GOOS=windows GOARCH=386 go build ${STATIC_FLAGS} -o bin/traffic-generator-windows-386.exe .
	env CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build ${STATIC_FLAGS} -o bin/traffic-generator-linux-amd64 .
	env CGO_ENABLED=0 GOOS=linux GOARCH=386 go build ${STATIC_FLAGS} -o bin/traffic-generator-linux-386 .