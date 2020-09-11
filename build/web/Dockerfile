FROM golang:1.15.2-alpine AS modcache
WORKDIR /workspace
COPY go.mod go.sum ./
RUN go mod download

FROM golang:1.15.2-alpine AS builder
WORKDIR /workspace
COPY --from=modcache /go/pkg /go/pkg
COPY . .
RUN go build -v -o /build/app ./cmd/web

FROM alpine:latest  
WORKDIR /root/
COPY --from=builder /build/app .
CMD ["./app"]  