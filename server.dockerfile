FROM golang:1.18
RUN go env -w GO111MODULE=on
RUN go env -w GOPROXY="https://goproxy.cn"
RUN mkdir /app
WORKDIR /app
ADD . /app/
RUN go mod tidy
RUN go build -o app main.go
RUN chmod +x ./app
ENTRYPOINT ["./app"]
