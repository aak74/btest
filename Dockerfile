FROM golang:1.21-alpine AS build

WORKDIR /app

COPY go.mod ./
COPY go.sum ./
RUN go mod download

COPY ./cmd ./cmd

RUN go build -o /btest ./cmd/btest

##
## Deploy
##
FROM alpine:3.15.0

WORKDIR /

COPY --from=build /btest /btest

EXPOSE 8080

ENTRYPOINT ["/btest"]