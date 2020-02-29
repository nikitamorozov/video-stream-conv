FROM alpine:3.11.3

MAINTAINER Nikita Morozov <morozov.nikita@me.com>

RUN mkdir /var/app/video -p

WORKDIR /var/app
COPY ./config.json /var/app
COPY ./video-stream-conv /var/app

EXPOSE 6754
ENTRYPOINT "./video-stream-conv"