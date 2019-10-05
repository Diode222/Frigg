FROM golang:1.13
MAINTAINER Diode "diodebupt@163.com"
WORKDIR $GOPATH/src/github.com/Diode222/Frigg
ADD . $GOPATH/src/github.com/Diode222/Frigg
ENV GO111MODULE on
ENV GOPROXY https://goproxy.io
RUN go mod download && go build main.go
EXPOSE 42222
ENTRYPOINT  ["./main"]
