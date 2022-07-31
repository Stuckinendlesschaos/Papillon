FROM golang:latest
WORKDIR /go/src
ENV GO111MODULE=on
ENV GOPROXY=https://goproxy.cn

ADD . /go/src
RUN go mod tidy
RUN go build
CMD ["/go/src/papillon"]
