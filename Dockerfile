FROM golang:1.15-alpine

WORKDIR /app
COPY . .

RUN apk update && \
    apk upgrade && \
    apk add --no-cache alpine-sdk build-base tzdata git zsh
