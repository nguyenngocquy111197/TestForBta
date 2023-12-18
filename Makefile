.DEFAULT_GOAL := start
.PHONY: clean test builder start
SHELL = /usr/bin/env bash

SERVICE_NAME := BOOKING-SERVICE
VPREFIX := app

MAIN_PATH := ./exec/pricing
DEFAULT_EXEC := ./bin/pricing.exec

MAIN_ADDR ?= ${ADDR}
ifeq (${MAIN_ADDR},)
	MAIN_ADDR := 0.0.0.0:8080                                                           
endif

VERSION ?= ${IMAGE_TAG}
ifeq (${VERSION},)
	VERSION := dev
endif

CGO_ENABLED := 0 # 0 or 1
GOOS := darwin # linux or windows or darwin 
GOARCH := amd64 # amd64 or arm
GIN_MODE := debug # debug or release


GIT_REVISION := $(shell git rev-parse --short HEAD)
GIT_BRANCH := $(shell git rev-parse --abbrev-ref HEAD)
GO_LDFLAGS   := -s -w -X ${VPREFIX}.Service=${SERVICE_NAME} -X ${VPREFIX}.Branch=${GIT_BRANCH} -X ${VPREFIX}.Version=${VERSION} -X ${VPREFIX}.Revision=${GIT_REVISION} -X ${VPREFIX}.BuildUser=$(shell whoami)@$(shell hostname) -X ${VPREFIX}.BuildDate=$(shell date -u +"%Y-%m-%dT%H:%M:%SZ")
GO_FLAGS     := -ldflags "-extldflags \"-static\" ${GO_LDFLAGS}" 

start: clean test builder
	@echo -ne "- Start service ...\r\n\n>>> Application running ...\r\n\n"
ifeq ($(o),)
	@GIN_MODE=${GIN_MODE} MAIN_ADDR=${MAIN_ADDR} MEDIA_ADDR=${MEDIA_ADDR} ${DEFAULT_EXEC}
else
	@GIN_MODE=${GIN_MODE} MAIN_ADDR=${MAIN_ADDR} MEDIA_ADDR=${MEDIA_ADDR} ./$(o) || $(o)
endif

builder: ${MAIN_PATH}/main.go
	@echo -ne "- Building the program from source code ...\r\n\n"
ifeq ($(o),)
	@export CGO_ENABLED=${CGO_ENABLED};\
		export GOOS=${GOOS};\
		export GOARCH=${GOARCH};\
		go build  ${GO_FLAGS} -o ${DEFAULT_EXEC} -trimpath ${MAIN_PATH}/*.go
	@chmod +x ${DEFAULT_EXEC}
else
	@export CGO_ENABLED=${CGO_ENABLED};\
		export GOOS=${GOOS};\
		export GOARCH=${GOARCH};\
		go build  ${GO_FLAGS} -o $(o) -trimpath ${MAIN_PATH}/*.go
	@chmod +x $(o)
endif
	@echo -ne "\n[OK]\n\n"

test:
	@echo -ne "- Try re-test again ...\r\n\n";\
		go test -v -count=1 ./...
	@echo -ne "\n[OK]\n\n"

clean:
	@echo -ne "- Clean up the build directory ..."
ifeq ($(o),)
	@rm -rf ./bin
else
	@rm -rf $(o)
endif
	@echo " [OK]"

 