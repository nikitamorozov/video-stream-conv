FROM golang:1.6-alpine

MAINTAINER Nikita Morozov <morozov.nikita@me.com>

RUN apk --no-cache add ca-certificates curl bash xz-libs git
WORKDIR /tmp
RUN curl -L -O http://johnvansickle.com/ffmpeg/releases/ffmpeg-release-64bit-static.tar.xz
RUN tar -xf ffmpeg-release-64bit-static.tar.xz && \
      cd ff* && mv ff* /usr/local/bin

WORKDIR /

ENTRYPOINT ["/bin/bash"]