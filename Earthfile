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
    COPY --dir cmd internal pkg api web space grid  .

    RUN --mount=type=cache,target=/root/.cache/go-build \
        go build -v -o ./.output/api ./api/cmd/api && \    
        go build -v -o ./.output/currentturn ./cmd/currentturn && \
        go build -v -o ./.output/gamerepo ./cmd/gamerepo && \
        go build -v -o ./.output/web ./web/cmd/web && \
        go build -v -o ./.output/grid ./grid/cmd/grid && \
        go build -v -o ./.output/checker ./cmd/checker && \
        go build -v -o ./.output/turncontroller ./cmd/turncontroller && \
        go build -v -o ./.output/space ./space/cmd/space
    
    SAVE ARTIFACT ./.output/* AS LOCAL ./.output/

protobuild:
    FROM +deps
    RUN apk add protoc
    RUN go get google.golang.org/grpc@v1.27.0
    RUN go get github.com/golang/protobuf/protoc-gen-go@v1.4.2
    RUN go install google.golang.org/grpc/cmd/protoc-gen-go-grpc
    SAVE IMAGE

protos:
    BUILD ./pkg/game/rpcrepository/+protos
    BUILD ./grid/pkg/grid/rpcgrid/+protos
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