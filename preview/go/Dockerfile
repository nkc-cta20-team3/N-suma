# build environment
FROM golang:1.19.1-alpine as build-stage
RUN apk update 
COPY ./go/src/main /go/src/main
COPY ./preview/go/n-suma-firebase-admin.secret /.select/n-suma-firebase-admin.secret
WORKDIR /go/src/main
RUN GOOS=linux go build -o ./build/main ./main.go

# production environment
FROM alpine:latest as production-stage
COPY ./preview/go/n-suma-firebase-admin.secret /.select/n-suma-firebase-admin.secret
COPY --from=build-stage /go/src/main/build/main /app/main
EXPOSE 8080
CMD ["/app/main"]


