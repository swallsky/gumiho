FROM golang:alpine3.15
LABEL maintainer="JackXu <xjz1688@163.com>"

USER root

# 更改镜像源为阿里云
RUN sed -i 's/dl-cdn.alpinelinux.org/mirrors.aliyun.com/' /etc/apk/repositories \
    && apk update \
    && apk upgrade \
    && apk add --no-cache bash

COPY ./build/ /opt/


WORKDIR /opt/

CMD ["./fox","server","up"]

EXPOSE 8081