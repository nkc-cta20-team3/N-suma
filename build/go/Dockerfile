FROM golang:1.19.1-alpine

WORKDIR /go/src/main

RUN apk update && go install github.com/cosmtrek/air@latest && go install github.com/swaggo/swag/cmd/swag@latest 

COPY ./n-suma-firebase-admin.secret /.select/n-suma-firebase-admin.secret

COPY ./run.sh /run.sh
RUN chmod +x /run.sh
RUN apk add --no-cache dos2unix
RUN dos2unix /run.sh
CMD [ "sh","/run.sh" ]