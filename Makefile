# Go parameters
GOCMD=go
GOBUILD=$(GOCMD) build
GORUN=$(GOCMD) run

#Binary name
BINARY_NAME=ped-consul-watch

#Main Files
CMD_PATH=./cmd/http-app/main.go

build: build-ped-consul-watch

build-ped-consul-watch:
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 ${GOBUILD} -ldflags='-s -w' -o ./bin/${BINARY_NAME}-linux -v ${CMD_PATH}
	GOOS=darwin GOARCH=amd64 ${GOBUILD} -o ./bin/${BINARY_NAME}-darwin -v ${CMD_PATH}

run-ped-consul-watch-linux: build-ped-consul-watch
	./bin/ped-consul-watch-linux

run-ped-consul-watch-darwin: build-ped-consul-watch
	./bin/ped-consul-watch-darwin

