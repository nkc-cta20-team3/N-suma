FROM golang:1.19.1-alpine

RUN apk update 

WORKDIR /go/src/

RUN go mod init main && \
    go get github.com/gin-gonic/gin
