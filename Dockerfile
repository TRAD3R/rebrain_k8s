FROM golang:1.23 as builder
COPY . /go/src/rebrain_k8s
WORKDIR /go/src/rebrain_k8s/cmd

RUN GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build -ldflags='-w -s' -o /go/bin/service

FROM alpine:latest
COPY --from=builder /go/bin/service /go/bin/service
ENTRYPOINT["/go/bin/service"]