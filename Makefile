_@=@

# Version related variables.
VERSION=$(shell $(GOBASE)/.tools/cmd/git-revision.sh)
BUILD_DATE=$(shell date -u +%Y/%m/%d-%H:%M:%S)
GIT_COMMIT=$(shell git rev-parse --verify HEAD)
PROJECTNAME=$(shell basename "$(PWD)")

# Go related variables.
GOBASE=$(shell pwd)
GOBIN=$(GOBASE)/bin
GOBINARY=$(GOBIN)/$(PROJECTNAME)
GOFILES=$(wildcard **/**/*.go) $(wildcard **/*.go) $(wildcard *.go)
GOMAIN=$(GOBASE)
PKGS=$(shell go list ./... | grep -v /vendor)
LDFLAGS=-ldflags "-X main.version=$(VERSION) -X main.buildDate=$(BUILD_DATE) -X main.gitCommit=$(GIT_COMMIT)"

RUN_CMD=$(GOBINARY)
TEST_CMD=go test

# Tools Configuration
NOTIFY_CMD=notify-send
NOTIFY_COMPILE_CMD=$(NOTIFY_CMD) -i $(GOBASE)/.tools/icons/gnu.png 'Makefile'

all: $(GOBINARY)
	$(_@) echo "build done"
.PHONY: all

build: $(GOBINARY)
.PHONY: build

test: $(GOFILES)
	$(_@) $(TEST_CMD) -v ./...
.PHONY: test

run: $(GOBINARY)
	$(_@) $(RUN_CMD)
.PHONY: run

$(GOBINARY): $(GOFILES)
	$(_@) go mod vendor &>/dev/null
	$(_@) go build $(LDFLAGS) -o $@ $(GOMAIN)
	$(_@) $(NOTIFY_COMPILE_CMD) '$(PROJECTNAME) $(VERSION): build completed' &>/dev/null || exit 0

init-git: $(GOBASE)/.tools/cmd/githooks/*
	$(_@) @for f in $(shell ls ${GOBASE}/.tools/cmd/githooks); do ln -s $(GOBASE)/.tools/cmd/githooks/$${f} $(GOBASE)/.git/hooks/; done
.PHONY: init-git

clean:
	$(_@) cd $(GOBIN) && ls | grep -v .gitkeep | xargs rm -rf && cd $(GOBASE) &>/dev/null && rm -r vendor &>/dev/null || exit 0
.PHONY: clean
