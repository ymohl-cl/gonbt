IGNORED_FOLDER=.ignore
MODULE_NAME := $(shell go list -m)

all: install lint build

.PHONY: install
install:
	@go mod download
	@go get github.com/golang/mock/mockgen@v1.5.0
	@go get -u golang.org/x/lint/golint

.PHONY: build
build:
	@go build ./...

test:
	@go test -count=1 ./...

.PHONY: lint
lint:
	golint ./...

.PHONY: mock
mock:
	@mockgen -source=reader.go -destination=mockreader_test.go -package=gonbt
	@mockgen -source=writer.go -destination=mockwriter_test.go -package=gonbt

.PHONY: clean
clean:
	@rm -rf ${IGNORED_FOLDER} 

.PHONY: fclean
fclean: clean
	@rm -rf ${BIN_FOLDER}
