define servicerule
.build/$1/app: generated
	@echo Building service $1
	CGO_ENABLED=0 GOOS=linux go build -o ./.build/$1/app -ldflags "-X github.com/theothertomelliott/tic-tac-toverengineered/common/version.Version=$(VERSION)"	./services/$1/cmd/$1

.build/$1/app_local: generated
	@echo Building service $1
	go build -o ./.build/$1/app_local -ldflags "-X github.com/theothertomelliott/tic-tac-toverengineered/common/version.Version=$(VERSION)"	./services/$1/cmd/$1

.PHONY: $1 $1_local
$1: .build/$1/app
$1_local: .build/$1/app_local
endef