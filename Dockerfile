FROM golang:1.18.1-bullseye

RUN apt-get update -y \
    && apt-get install -y git

WORKDIR /whitetail-build
COPY whitetail /whitetail-build
RUN env GO111MODULE=on GOOS=linux CGO_ENABLED=0 go build -v -o whitetail

FROM ubuntu:latest

USER 0

RUN adduser --disabled-password whitetail

WORKDIR /whitetail

COPY --from=0 /whitetail-build/whitetail ./

RUN apt-get update \
    && apt-get install -y \
        bash \
        python3

SHELL ["/bin/bash", "-c"]

# Copy over built UI files
COPY whitetail/ui-dist /whitetail

# make it executable
RUN chmod +x /whitetail/whitetail

# RUN chown -R whitetail:whitetail /whitetail

CMD ["/whitetail/whitetail"]
