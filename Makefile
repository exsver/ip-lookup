GO_BIN ?= go
OUT_BIN = ip-lookup

export PATH := $(PATH):/usr/local/go/bin

all: clean build

build:
	$(GO_BIN) mod tidy
	$(GO_BIN) build -o $(OUT_BIN) -v

download:
	$(GO_BIN) get
	$(GO_BIN) mod tidy

update:
	$(GO_BIN) get -u
	$(GO_BIN) mod tidy

clean:
	$(GO_BIN) clean
	rm -f $(OUT_BIN)

test:
	$(GO_BIN) test -failfast ./...

lint:
	# https://golangci-lint.run/welcome/install/#local-installation
	# linter binary will be $(go env GOPATH)/bin/golangci-lint
	#curl -sSfL https://raw.githubusercontent.com/golangci/golangci-lint/master/install.sh | sh -s -- -b $(go env GOPATH)/bin v1.61.0
	golangci-lint run ./...