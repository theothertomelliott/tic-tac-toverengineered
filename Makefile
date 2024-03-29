all: generated

include generated.mk
include files.mk

PROTOFILES := $(call rwildcard,./,*.proto)
PBFILES := $(subst .proto,.pb.go,$(PROTOFILES))
GRPCPBFILES := $(subst .proto,_grpc.pb.go,$(PROTOFILES))
VERSION ?= dev

# Create rules to build protos
$(foreach proto,$(PROTOFILES),$(eval $(call protorule,$(proto))))

protos: installgoproto $(PBFILES) $(GRPCPBFILES)

include services.mk

SERVICEDIRS = $(wildcard services/*)
SERVICES = $(subst services/,,$(SERVICEDIRS))
GEN_GO = $(rwildcard ./,*.gen.go)

# Create rules to build services
$(foreach service,$(SERVICES),$(eval $(call servicerule,$(service))))

services: $(SERVICES)
services_local: $(addsuffix _local,$(SERVICES))

include docker.mk

$(foreach service,$(SERVICES),$(eval $(call dockerrule,$(service))))

docker: services $(addsuffix _docker,$(SERVICES))
docker_push: $(addsuffix _docker_push,$(SERVICES))

include web.mk

$(eval $(call tailwind,web))
$(eval $(call views,web))
$(eval $(call js,web))

webdeps: web_views web_tailwind web_js
web: webdeps
web_local: webdeps

.PHONY: generated
generated: protos openapi

test: generated
	go test ./...

testcover: generated
	go test -coverprofile=coverage.out ./...

testshort: generated
	go test -short ./...

.PHONY: clean protos test testshort services web docker docker_push
clean:
	rm -rf .bin
	rm -rf .build
	rm -f $(PBFILES) $(GRPCPBFILES)
	rm -f $(GEN_GO)
	rm -f coverage.out