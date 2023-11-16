FROM golang:1.19.1-alpine

WORKDIR /go/src/main

RUN apk update && go install github.com/cosmtrek/air@latest && go install github.com/swaggo/swag/cmd/swag@latest 

COPY ./n-suma-firebase-admin.secret /go/select/n-suma-firebase-admin.secret

CMD [ "air", "-c", ".air.toml" ]
