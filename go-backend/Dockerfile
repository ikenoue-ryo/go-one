FROM golang:1.16.3-buster

WORKDIR /usr/local/go/src/go-one/go-backend

ADD . /usr/local/go/src/go-one/go-backend

RUN go get gopkg.in/ini.v1
RUN go mod tidy
RUN go get github.com/cosmtrek/air
CMD ["air"]
