FROM golang:alpine
LABEL maintainer="shooter<byshooter@163.com>"
ENV GO111MODULE=on \
    CGO_ENABLED=0 \
    GOOS=linux \
    GOARCH=amd64 \
    GOPROXY="https://mirrors.aliyun.com/goproxy/,direct"
WORKDIR /app
COPY . .
RUN go version && go build -o anymsg ./main

FROM scratch
COPY --from=0 /app/cfg.json /anymsg/config/cfg.json
COPY --from=0 /app/anymsg /anymsg/anymsg
EXPOSE 4000
ENTRYPOINT ["/anymsg/anymsg","--config=/anymsg/config/cfg.json"]
