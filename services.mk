define servicerule
.build/$1/app: protos
	@echo Building service $1
	CGO_ENABLED=0 GOOS=linux go build -o ./.build/$1/app -ldflags "-X github.com/theothertomelliott/tic-tac-toverengineered/common/version.Version=$(VERSION)"	./services/$1/cmd/$1

.PHONY: $1
$1: .build/$1/app
endef