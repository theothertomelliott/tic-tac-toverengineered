# syntax = docker/dockerfile:1-experimental

FROM golang:1.15.2-alpine AS modcache
WORKDIR /workspace
COPY go.mod go.sum ./
RUN go mod download

FROM golang:1.15.2-alpine AS builder
WORKDIR /workspace
COPY --from=modcache /go/pkg /go/pkg
COPY . .
RUN --mount=type=cache,target=/root/.cache/go-build \
    go build -v -o /build/app ./cmd/gamerepo

FROM alpine:latest  
WORKDIR /root/
COPY --from=builder /build/app .
CMD ["./app"]  