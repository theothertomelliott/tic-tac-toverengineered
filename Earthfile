FROM golang:1.15-alpine
WORKDIR /go-example

deps:
    COPY go.mod go.sum ./
    RUN go mod download
    SAVE ARTIFACT go.mod AS LOCAL go.mod
    SAVE ARTIFACT go.sum AS LOCAL go.sum
    SAVE IMAGE

binaries:
    FROM +deps
    COPY . .

    # Api
    RUN --mount=type=cache,target=/root/.cache/go-build \
        go build -v -o api ./cmd/api
    SAVE ARTIFACT api api AS LOCAL .output/api
    
    # Current turn
    RUN --mount=type=cache,target=/root/.cache/go-build \
        go build -v -o currentturn ./cmd/currentturn
    SAVE ARTIFACT currentturn currentturn AS LOCAL .output/currentturn 
    
    # Game repo
    RUN --mount=type=cache,target=/root/.cache/go-build \
        go build -v -o gamerepo ./cmd/gamerepo
    SAVE ARTIFACT gamerepo gamerepo AS LOCAL .output/gamerepo
    
    # Web
    RUN --mount=type=cache,target=/root/.cache/go-build \
        go build -v -o web ./cmd/web
    SAVE ARTIFACT web web AS LOCAL .output/web

    # Grid
    RUN --mount=type=cache,target=/root/.cache/go-build \
        go build -v -o grid ./cmd/grid
    SAVE ARTIFACT grid grid AS LOCAL .output/grid

    # Win Checker
    RUN --mount=type=cache,target=/root/.cache/go-build \
        go build -v -o checker ./cmd/checker
    SAVE ARTIFACT checker checker AS LOCAL .output/checker

    SAVE IMAGE

protobuild:
    FROM +deps
    RUN apk add protoc
    RUN go get google.golang.org/grpc@v1.27.0
    RUN go get github.com/golang/protobuf/protoc-gen-go@v1.4.2
    RUN go install google.golang.org/grpc/cmd/protoc-gen-go-grpc
    SAVE IMAGE

protos:
    BUILD ./pkg/game/rpcrepository/+protos
    BUILD ./pkg/grid/rpcgrid/+protos
    BUILD ./pkg/turn/rpcturn/+protos
    BUILD ./pkg/win/rpcchecker/+protos

images:
    BUILD ./build/web+docker
    BUILD ./build/grid+docker
    BUILD ./build/api+docker
    BUILD ./build/currentturn+docker
    BUILD ./build/gamerepo+docker
    BUILD ./build/checker+docker

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
    COPY ./build/start.sh .
    COPY ./build/restart.sh .
    ENTRYPOINT ["./start.sh", "/root/app"]
    ARG VERSION=dev
    ARG IMAGE_REF=api-image:$VERSION
    RUN echo "Building image with tag $IMAGE_REF"
    SAVE IMAGE $IMAGE_REF