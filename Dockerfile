FROM jrottenberg/ffmpeg:4.1-ubuntu

MAINTAINER Nikita Morozov <morozov.nikita@me.com>

RUN apt-get install -y bc

RUN mkdir /var/app/video -p

WORKDIR /var/app
COPY ./config.json /var/app
COPY ./video-stream-conv /var/app
COPY converter.sh /var/app

EXPOSE 6754
ENTRYPOINT "./video-stream-conv"