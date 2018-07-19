.PHONY: build test clean codegen dockerize
OUTPUT = ./riff-init
TAG ?= $(shell cat VERSION)
GO_SOURCES = $(shell find cmd pkg -type f -name '*.go' -not -name 'mock_*.go')

build: $(OUTPUT)

test: build
	go test ./...

$(OUTPUT): $(GO_SOURCES)
	go build -o $(OUTPUT) cmd/main.go

clean:
	rm -f $(OUTPUT)

vendor:
	./hack/update-deps.sh

codegen:
	./hack/update-codegen.sh

dockerize: .dockerize

.dockerize: $(GO_SOURCES) Dockerfile
	docker build . -t projectriff/riff-init:$(TAG) && touch .dockerize