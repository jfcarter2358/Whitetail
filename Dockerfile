FROM golang:1.18.1-alpine

RUN apk update && apk add git

WORKDIR /whitetail-build
COPY whitetail /whitetail-build
RUN env GO111MODULE=on GOOS=linux CGO_ENABLED=0 go build -v -o whitetail

FROM alpine:latest

USER 0

RUN adduser --disabled-password whitetail

WORKDIR /whitetail

COPY --from=0 /whitetail-build/whitetail ./

RUN apk update \
    && apk add \
    bash

SHELL ["/bin/bash", "-c"]

RUN mkdir /whitetail/data
RUN mkdir /whitetail/saved
RUN mkdir -p /whitetail/config/custom/icon
RUN mkdir -p /whitetail/config/custom/logo

# add resources
COPY resources/static /whitetail/static
COPY resources/templates /whitetail/templates

# make it executable
RUN chmod +x /whitetail/whitetail

RUN chown -R whitetail:whitetail /whitetail

USER whitetail

CMD ["/whitetail/whitetail"]
