SPEC_URL ?= https://docs.devin.ai/openapi.yaml
SPEC_FILE := openapi/openapi.yaml

.PHONY: all fetch-spec generate build test tidy clean

all: generate build

## fetch-spec: download the latest OpenAPI spec
fetch-spec:
	@mkdir -p openapi
	curl -fsSL "$(SPEC_URL)" -o "$(SPEC_FILE)"
	@echo "Fetched $(SPEC_FILE) ($$(wc -l < $(SPEC_FILE)) lines)"

## generate: regenerate the Go client from the local spec
generate:
	go generate ./...
	go mod tidy

## build: compile the SDK
build:
	go build ./...

## test: run tests
test:
	go test ./...

## tidy: sync go.mod / go.sum
tidy:
	go mod tidy

## clean: remove generated code
clean:
	rm -f api/devin.gen.go
