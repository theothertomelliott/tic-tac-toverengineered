GOINSTALL ?= go install

GOBIN:=$(CURDIR)/.bin/
PATH:="$(GOBIN):$(PATH)"

# Always go install to the bin dir
GOINSTALL := GOBIN=$(GOBIN) $(GOINSTALL)

## Protos

# Use gobin for protoc plugins
PROTOC ?= protoc
PROTOC := PATH=$(PATH) $(PROTOC)

.bin/protoc-*:
	$(GOINSTALL) google.golang.org/protobuf/cmd/protoc-gen-go; \
	$(GOINSTALL) google.golang.org/grpc/cmd/protoc-gen-go-grpc

installgoproto: .bin/protoc-*

define protorule
$(subst .proto,.pb.go,$1) $(subst .proto,_grpc.pb.go,$1): $1
	$(call buildproto,$1)
endef

buildproto = $(PROTOC) \
	--go_out=. \
	--go-grpc_out=. \
    --go_opt=paths=source_relative \
	--go-grpc_opt=paths=source_relative \
    $1

# OpenAPI

OPENAPISRC := services/api/pkg/tictactoeapi/tictactoe.openapi.yaml

## OpenAPI Go Codegen

# Use gobin for OpenAPI generation
OPENAPIGENGO := ./.bin/oapi-codegen

.bin/oapi-codegen:
	$(GOINSTALL) github.com/deepmap/oapi-codegen/cmd/oapi-codegen

services/api/pkg/tictactoeapi/tictactoe.gen.go: $(OPENAPIGENGO) $(OPENAPISRC)
	$(OPENAPIGENGO) -package tictactoeapi $(OPENAPISRC)  > services/api/pkg/tictactoeapi/tictactoe.gen.go

openapi: services/api/pkg/tictactoeapi/tictactoe.gen.go

## OpenAPI Generator (for JavaScript)

OPENAPIGENERATOR := ./.bin/openapi-generator-cli.jar
WGET := wget
JAVA := java
JSCLIENTDIR := frontend/js/tictactoe-client

$(OPENAPIGENERATOR):
	$(WGET) https://repo1.maven.org/maven2/org/openapitools/openapi-generator-cli/5.2.1/openapi-generator-cli-5.2.1.jar -O $(OPENAPIGENERATOR)

getopenapigen: $(OPENAPIGENERATOR)

$(JSCLIENTDIR): $(OPENAPISRC)
	$(JAVA) -jar $(OPENAPIGENERATOR) generate -g javascript -i $(OPENAPISRC) -o $(JSCLIENTDIR)

openapijsclient: getopenapigen frontend/js/tictactoe-client