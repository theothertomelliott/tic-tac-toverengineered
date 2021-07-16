define dockerrule
$1_docker: dockerbaseimage $1
	@echo Building docker image for $1
	docker build -f docker/app/Dockerfile -t "docker.io/tictactoverengineered/$1:$(VERSION)" .build/$1

$1_docker_push:
	docker image push docker.io/tictactoverengineered/$1:$(VERSION)

.PHONY: $1_docker $1_docker_push
endef

dockerbaseimage:
	docker build -t "docker.io/tictactoverengineered/base:latest" docker/base