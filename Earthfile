FROM golang:1.15-alpine
WORKDIR /go-example

deps:
    COPY go.mod go.sum ./
    RUN go mod download
    SAVE ARTIFACT go.mod AS LOCAL go.mod
    SAVE ARTIFACT go.sum AS LOCAL go.sum
    SAVE IMAGE

source:
    FROM +deps
    COPY . .
    SAVE IMAGE

binaries:
    FROM +source
    RUN --mount=type=cache,target=/root/.cache/go-build \
        go build -v -o api ./cmd/api
    SAVE ARTIFACT api api AS LOCAL .output/api
    RUN --mount=type=cache,target=/root/.cache/go-build \
        go build -v -o currentturn ./cmd/currentturn
    SAVE ARTIFACT currentturn currentturn AS LOCAL .output/currentturn 
    RUN --mount=type=cache,target=/root/.cache/go-build \
        go build -v -o gamerepo ./cmd/gamerepo
    SAVE ARTIFACT gamerepo gamerepo AS LOCAL .output/gamerepo
    RUN --mount=type=cache,target=/root/.cache/go-build \
        go build -v -o web ./cmd/web
    SAVE ARTIFACT web web AS LOCAL .output/web

protobuild:
    FROM +deps
    RUN apk add protoc
    RUN go get google.golang.org/grpc@v1.27.0
    RUN go get github.com/golang/protobuf/protoc-gen-go@v1.4.2
    RUN go install google.golang.org/grpc/cmd/protoc-gen-go-grpc
    SAVE IMAGE

protos:
    BUILD ./pkg/game/rpcrepository/+protos
    BUILD ./pkg/turn/rpcturn/+protos

all:
    BUILD ./build/web+build
    BUILD ./build/api+build
    BUILD ./build/currentturn+build
    BUILD ./build/gamerepo+build

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