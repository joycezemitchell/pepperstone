VERSION=v1.0.0
TODAY = $(shell date +"%Y%m%d_%H%M%S")
LDFLAGS=-ldflags " -X main.version=${VERSION}.build.${TODAY}"
BUILDFLAGS=
NAME=bin/pepperstonecsv
MAIN=main.go


all: app

build:
	go build ${BUILDFLAGS} ${LDFLAGS} -o ${NAME} ${MAIN}

linux:
	GO111MODULE=on GOPRIVATE=github.com/PaceNow/* CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -a -installsuffix cgo -o ${NAME} ${BUILDFLAGS} ${LDFLAGS} ${MAIN}

run:
	go run main.go