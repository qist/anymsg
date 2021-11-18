FROM golang:alpine
LABEL maintainer="shooter<byshooter@163.com>"

ENV GO111MODULE=on \
    CGO_ENABLED=0 \
    GOOS=linux \
    GOARCH=amd64 \
    GOPROXY="https://mirrors.aliyun.com/goproxy/,direct"
	
WORKDIR /app

COPY . .
RUN go version
RUN go build -o anymsg .

EXPOSE 4000

ENTRYPOINT ["./anymsg"]