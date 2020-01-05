FROM golang:1.13.4-alpine

ENV LANG=ja_JP.UTF-8

RUN apk update \
    && apk upgrade \
    && apk add --update --no-cache \
       git \
       gcc \
       libc-dev \
    && rm -rf /var/cache/apk/*
#RUN go get github.com/mattn/go-sqlite3

RUN mkdir -p /native_message
WORKDIR /native_message

ENTRYPOINT tail -f /dev/null
