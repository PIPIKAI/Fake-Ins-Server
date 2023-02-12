FROM golang:1.18
RUN go env -w GO111MODULE=on
RUN go env -w GOPROXY="https://goproxy.cn"
RUN mkdir /app
WORKDIR /app
ADD  . /app
RUN go mod tidy
RUN  CGO_ENABLED=0 go build -a -o app -ldflags '-extldflags "-static"' .
RUN chmod +x ./app
ENTRYPOINT ["./app"]
