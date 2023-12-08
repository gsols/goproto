SHELL=/bin/bash -e -o pipefail
PWD = $(shell pwd)

download: ## Downloads the dependencies
	@go mod download

tidy: ## Cleans up go.mod and go.sum
	@go mod tidy


# Go dependencies versioned through tools.go
GO_DEPENDENCIES = google.golang.org/protobuf/cmd/protoc-gen-go \
                  google.golang.org/grpc/cmd/protoc-gen-go-grpc \
                  github.com/envoyproxy/protoc-gen-validate \
                  github.com/bufbuild/buf/cmd/buf \
                  github.com/bufbuild/buf/cmd/protoc-gen-buf-breaking \
                  github.com/bufbuild/buf/cmd/protoc-gen-buf-lint\
                  golang.org/x/tools/cmd/goimports

define make-go-dependency
  # target template for go tools, can be referenced e.g. via /bin/<tool>
  bin/$(notdir $1):
	GOBIN=$(PWD)/bin go install $1
endef

# this creates a target for each go dependency to be referenced in other targets
$(foreach dep, $(GO_DEPENDENCIES), $(eval $(call make-go-dependency, $(dep))))

.PHONY: proto/buf.lock
proto/buf.lock: bin/buf
	@bin/buf mod update proto

protolint: proto/buf.lock bin/protoc-gen-buf-lint ## Lints your protobuf files
	bin/buf lint

#protobreaking: proto/buf.lock bin/protoc-gen-buf-breaking ## Compares your current protobuf with the version on master to find breaking changes
#	bin/buf breaking --against '.git#branch=main'

proto: ## Generates code from protobuf files
proto:  proto/buf.lock bin/protoc-gen-go bin/protoc-gen-go-grpc bin/protoc-gen-validate
	PATH=$(PWD)/bin:$$PATH buf generate


help: ## Shows the help
	@echo 'Usage: make <OPTIONS> ... <TARGETS>'
	@echo ''
	@echo 'Available targets are:'
	@echo ''
	@grep -E '^[ a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | \
        awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-20s\033[0m %s\n", $$1, $$2}'
	@echo ''
