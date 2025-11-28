.PHONY: i init
.PHONY: s serv
.PHONY: c client
.PHONY: r run
.PHONY: a all
.PHONY: api
.DEFAULT_GOAL := build

init:
	git submodule update --init --recursive
	go mod tidy

docker:
	docker build -t "gin-template" .

docker-run:
	docker run -d --restart always -p 127.0.0.1:8080:8080 "gin-template"

build:
	bash build.sh

build-debug:
	go build -v -gcflags="all=-N -l" -o out/serv cmd/serv/main.go

run:
	./out/serv

ifeq (controller,$(firstword $(MAKECMDGOALS)))
  CONTROLLER_ARGS := $(wordlist 2,$(words $(MAKECMDGOALS)),$(MAKECMDGOALS))
  $(eval $(CONTROLLER_ARGS):;@:)
endif

controller:
	@go run cmd/createController/main.go $(CONTROLLER_ARGS)

.PHONY: docker docker-run build build-debug run controller

serv: build
br: build run

rs: run
i: init
s: serv
c: client
r: run
a: all
all: build run