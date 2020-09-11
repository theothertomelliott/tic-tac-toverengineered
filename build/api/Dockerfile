FROM golang:1.15.1 AS builder
WORKDIR /go/src/github.com/theothertomelliott/tic-tac-toeverengineered
COPY cmd ./cmd
COPY internal ./internal
COPY pkg ./pkg
COPY go.mod .
COPY go.sum .
RUN ls
RUN CGO_ENABLED=0 GOOS=linux go build -a -o app ./cmd/api

FROM alpine:latest  
WORKDIR /root/
COPY --from=builder /go/src/github.com/theothertomelliott/tic-tac-toeverengineered/app .
CMD ["./app"]  