all: protos

include protos.mk
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

# Create rules to build services
$(foreach service,$(SERVICES),$(eval $(call servicerule,$(service))))

services: $(SERVICES)

include docker.mk

$(foreach service,$(SERVICES),$(eval $(call dockerrule,$(service))))

docker: services dockerbaseimage $(addsuffix _docker,$(SERVICES))
docker_push: $(addsuffix _docker_push,$(SERVICES))

include web.mk

$(eval $(call tailwind,web))
$(eval $(call views,web))

web: web_views
web: web_tailwind

test: protos
	go test ./...

testshort: protos
	go test -short ./...

.PHONY: clean protos test testshort services web docker docker_push
clean:
	rm -rf .bin
	rm -rf .build
	rm $(PBFILES) $(GRPCPBFILES)