FROM ubuntu:18.04

RUN apt-get update
RUN mkdir /whitetail
RUN mkdir /whitetail/data
RUN mkdir /whitetail/saved
RUN mkdir -p /whitetail/config/custom/icon
RUN mkdir -p /whitetail/config/custom/logo

# add whitetail distribution
ADD dist/whitetail /whitetail/whitetail

# add resources
ADD dist/config /whitetail/config
ADD dist/static /whitetail/static
ADD dist/templates /whitetail/templates

# make it executable
RUN chmod +x /whitetail/whitetail

WORKDIR /whitetail

CMD ["/whitetail/whitetail"]
