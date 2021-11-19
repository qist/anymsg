FROM golang:alpine
LABEL maintainer="shooter<byshooter@163.com>"
WORKDIR /home
RUN mkdir /lib64 && ln -s /lib/libc.musl-x86_64.so.1 /lib64/ld-linux-x86-64.so.2
COPY ./cfg.json .
COPY ./bin/anymsg /usr/bin/

EXPOSE 4000

ENTRYPOINT ["anymsg"]
