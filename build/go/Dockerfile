FROM golang:1.19.1-alpine

WORKDIR /go/src/main

RUN apk update && go install github.com/cosmtrek/air@latest && go install google.golang.org/api/option@latest && go install firebase.google.com/go/v4@latest && go install github.com/swaggo/swag/cmd/swag@latest 

COPY ./n-suma-firebase-admin.secret /go/secret/n-suma-firebase-admin.secret

CMD [ "air", "-c", ".air.toml" ]
