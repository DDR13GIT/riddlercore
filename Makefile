export APP=riddlercore

.PHONY: all test coverage
all: get build install
format:
	gofmt -l -s -w .
get:
	go get ./...
build:
	go build ./...
install:
	go install ./...
build-run:
	go build -v .
	./riddlercore serve -v
test:
	go test ./... -v -coverprofile .coverage.txt
	go tool cover -func .coverage.txt
coverage: test
	go tool cover -html=.coverage.txt
init:
	chmod +x install-golangci.sh
	./install-golangci.sh
	cp docker-compose.example.yml docker-compose.yml
	cp config.example.yml config.yml
	export GOSUMDB=off
	go mod tidy
	go mod vendor -v
	go build -v .
	git init
	git branch -m main
	git add . && git commit  -m "init: this is where it all begins..." -m "This project initialized by adam project. See more at: https://magic.pathao.com/common/adam"
	git config core.hooksPath .githooks

build-run-worker:
	go build -v .
	./riddlercore worker
