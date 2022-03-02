FROM golang:alpine3.15
LABEL maintainer="JackXu <xjz1688@163.com>"

USER root

# 更改镜像源为阿里云
RUN sed -i 's/dl-cdn.alpinelinux.org/mirrors.aliyun.com/' /etc/apk/repositories \
    && apk update \
    && apk upgrade \
    && apk add --no-cache bash

COPY ./build/ /opt/

# 更改代码可执行权限
RUN chmod +x /opt/script/*.sh

WORKDIR /opt/

CMD ["./tail","server","up"]

EXPOSE 8081