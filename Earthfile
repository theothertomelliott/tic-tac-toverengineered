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
    RUN go test ./...

docker:
    FROM alpine
    WORKDIR /root
    ARG BINARY
    COPY ./.output/$BINARY ./app
    COPY ./common/docker/entrypoint/start.sh .
    COPY ./common/docker/entrypoint/restart.sh .
    ENTRYPOINT ["./start.sh", "/root/app"]
    ARG VERSION=dev
    ARG IMAGE_REF=api-image:$VERSION
    RUN echo "Building image with tag $IMAGE_REF"
    SAVE IMAGE $IMAGE_REF