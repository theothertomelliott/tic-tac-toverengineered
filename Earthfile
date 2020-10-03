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
    COPY --dir cmd ./cmd
    COPY --dir internal ./internal
    COPY --dir pkg ./pkg
    COPY --dir api ./api
    COPY --dir web ./web
    COPY --dir space ./space

    # Api
    RUN --mount=type=cache,target=/root/.cache/go-build \
        go build -v -o ./.output/api ./api/cmd/api
    SAVE ARTIFACT ./.output/api api AS LOCAL .output/api
    
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
        go build -v -o ./.output/web ./web/cmd/web
    SAVE ARTIFACT ./.output/web web AS LOCAL .output/web

    # Grid
    RUN --mount=type=cache,target=/root/.cache/go-build \
        go build -v -o grid ./cmd/grid
    SAVE ARTIFACT grid grid AS LOCAL .output/grid

    # Win Checker
    RUN --mount=type=cache,target=/root/.cache/go-build \
        go build -v -o checker ./cmd/checker
    SAVE ARTIFACT checker checker AS LOCAL .output/checker

    # Turn Controller
    RUN --mount=type=cache,target=/root/.cache/go-build \
        go build -v -o turncontroller ./cmd/turncontroller
    SAVE ARTIFACT turncontroller turncontroller AS LOCAL .output/turncontroller

    # Space
    RUN --mount=type=cache,target=/root/.cache/go-build \
        go build -v -o ./.output/space ./space/cmd/space
    SAVE ARTIFACT ./.output/space space AS LOCAL .output/space

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
    BUILD ./space/pkg/rpcspace/+protos

images:
    BUILD ./web/build+docker
    BUILD ./build/grid+docker
    BUILD ./api/build+docker
    BUILD ./build/currentturn+docker
    BUILD ./build/gamerepo+docker
    BUILD ./build/checker+docker
    BUILD ./build/turncontroller+docker
    BUILD ./space/build+docker

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