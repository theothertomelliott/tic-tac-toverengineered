FROM golang:1.15-alpine
WORKDIR /go-example

deps:
    COPY go.mod go.sum ./
    RUN go mod download
    SAVE ARTIFACT go.mod AS LOCAL go.mod
    SAVE ARTIFACT go.sum AS LOCAL go.sum
    SAVE IMAGE

binarybuild:
    FROM +deps
    ARG BINARY
    ARG VERSION=dev
    COPY --dir common bot api web space grid checker currentturn gamerepo turncontroller .
    RUN --mount=type=cache,target=/root/.cache/go-build \
        go build \
        -ldflags "-X github.com/theothertomelliott/tic-tac-toverengineered/common/version.Version=$VERSION" \
        -o ./.output/$BINARY ./$BINARY/cmd/$BINARY
    SAVE ARTIFACT ./.output/$BINARY AS LOCAL ./.output/$BINARY

protobuild:
    FROM +deps
    RUN apk add protoc
    RUN go get google.golang.org/grpc@v1.27.0
    RUN go get github.com/golang/protobuf/protoc-gen-go@v1.4.2
    RUN go install google.golang.org/grpc/cmd/protoc-gen-go-grpc
    SAVE IMAGE

protos:
    BUILD ./gamerepo/pkg/game/rpcrepository/+protos
    BUILD ./grid/pkg/grid/rpcgrid/+protos
    BUILD ./currentturn/pkg/turn/rpcturn/+protos
    BUILD ./checker/pkg/win/rpcchecker/+protos
    BUILD ./space/pkg/rpcspace/+protos

testdeps:
    FROM golang:1.15
    WORKDIR /root
    COPY go.mod go.sum ./
    RUN go mod download
    SAVE IMAGE

test:
    FROM +testdeps
    COPY . .
    RUN go test -short ./...

docker:
    FROM alpine
    WORKDIR /root
    ARG BINARY
    COPY ./.output/$BINARY ./app
    COPY ./common/docker/entrypoint/start.sh .
    COPY ./common/docker/entrypoint/restart.sh .
    ENTRYPOINT ["./start.sh", "/root/app"]
    ARG VERSION=dev
    ARG IMAGE_REF
    RUN echo "Building image with tag $IMAGE_REF"
    SAVE IMAGE --push $IMAGE_REF

images:
    ARG REGISTRY=docker.io/tictactoverengineered
    ARG VERSION=dev
    BUILD --build-arg REGISTRY=$REGISTRY --build-arg VERSION=$VERSION ./api/build/+docker
    BUILD --build-arg REGISTRY=$REGISTRY --build-arg VERSION=$VERSION ./bot/build/+docker
    BUILD --build-arg REGISTRY=$REGISTRY --build-arg VERSION=$VERSION ./checker/build/+docker
    BUILD --build-arg REGISTRY=$REGISTRY --build-arg VERSION=$VERSION ./currentturn/build/+docker
    BUILD --build-arg REGISTRY=$REGISTRY --build-arg VERSION=$VERSION ./gamerepo/build/+docker
    BUILD --build-arg REGISTRY=$REGISTRY --build-arg VERSION=$VERSION ./grid/build/+docker
    BUILD --build-arg REGISTRY=$REGISTRY --build-arg VERSION=$VERSION ./space/build/+docker
    BUILD --build-arg REGISTRY=$REGISTRY --build-arg VERSION=$VERSION ./turncontroller/build/+docker
    BUILD --build-arg REGISTRY=$REGISTRY --build-arg VERSION=$VERSION ./web/build/+docker

allbinaries:
    BUILD +protos
    BUILD --build-arg BINARY=api +binarybuild
    BUILD --build-arg BINARY=bot +binarybuild
    BUILD --build-arg BINARY=checker +binarybuild
    BUILD --build-arg BINARY=currentturn +binarybuild
    BUILD --build-arg BINARY=gamerepo +binarybuild
    BUILD --build-arg BINARY=grid +binarybuild
    BUILD --build-arg BINARY=space +binarybuild
    BUILD --build-arg BINARY=turncontroller +binarybuild
    BUILD --build-arg BINARY=web +binarybuild