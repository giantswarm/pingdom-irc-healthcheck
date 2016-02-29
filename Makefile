PROJECT=pingdom-irc-healthcheck

BUILD_PATH := $(shell pwd)/.gobuild
GS_PATH := $(BUILD_PATH)/src/github.com/giantswarm
GOPATH := $(BUILD_PATH)

BIN := $(PROJECT)

VERSION := $(shell cat VERSION)
COMMIT := $(shell git rev-parse --short HEAD)

.PHONY: all clean docker

SOURCE=$(shell find . -name '*.go')

all: $(BIN)

clean:
	rm -rf $(BUILD_PATH) $(BIN)

.gobuild:
	@mkdir -p $(GS_PATH)
	@rm -f $(GS_PATH)/$(PROJECT) && cd "$(GS_PATH)" && ln -s ../../../.. $(PROJECT)

	@builder get dep https://github.com/thoj/go-ircevent $(GOPATH)/src/github.com/thoj/go-ircevent
	@builder get dep https://github.com/juju/errgo $(GOPATH)/src/github.com/juju/errgo

$(BIN): $(SOURCE) VERSION .gobuild
	@echo Building inside Docker container for $(GOOS)/$(GOARCH)
	docker run \
	    --rm \
	    -v $(shell pwd):/usr/code \
	    -e GOPATH=/usr/code/.gobuild \
	    -e GOOS=$(GOOS) \
	    -e GOARCH=$(GOARCH) \
	    -w /usr/code \
	    golang:1.6 \
	    go build -a -o $(BIN)

docker: $(BIN)
	docker build -t $(PROJECT):latest .