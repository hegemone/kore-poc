REGISTRY         ?= quay.io
ORG              ?= hegemone
TAG              ?= latest
IMAGE            ?= $(REGISTRY)/$(ORG)/korecomm-go-poc
BUILD_DIR        = "${GOPATH}/src/github.com/hegemone/kore-poc/korecomm-go/build"
SOURCES          := $(shell find . -name '*.go' -not -path "*/vendor/*" -not -path "*/extensions/*")
PROJECT_ROOT := ""$(abspath $(lastword $(MAKEFILE_LIST)))/..""
.DEFAULT_GOAL    := build

vendor:
	@glide install -v

plugins: $(PLUGIN_SOURCES)
	@go build -buildmode=plugin -o ${BUILD_DIR}/bacon.plugins.kore.nsk.io.so -i -ldflags="-s -w" ./pkg/extension/plugin/bacon.go

adapters: $(ADAPTER_SOURCES)
	@go build -buildmode=plugin -o ${BUILD_DIR}/ex-discord.adapters.kore.nsk.io.so -i -ldflags="-s -w" ./pkg/extension/adapter/discord.go
	@go build -buildmode=plugin -o ${BUILD_DIR}/ex-irc.adapters.kore.nsk.io.so -i -ldflags="-s -w" ./pkg/extension/adapter/irc.go

korecomm: $(SOURCES) adapters plugins
	@go build -o ${BUILD_DIR}/korecomm -i -ldflags="-s -w" ./cmd/korecomm

build: korecomm
	@echo > /dev/null

clean:
	@rm -rf ${BUILD_DIR}

run: korecomm
	@KORECOMM_PLUGIN_DIR=${PROJECT_ROOT}/build \
	KORECOMM_ADAPTER_DIR=${PROJECT_ROOT}/build \
	./build/korecomm

image:
	docker build -t ${IMAGE}:${TAG} ${PROJECT_ROOT}

run-image:
	docker run -it ${IMAGE}:${TAG}

push:
	docker push ${IMAGE}:${TAG}

.PHONY: vendor image push clean run image run-image push
