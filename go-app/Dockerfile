FROM golang:latest

WORKDIR /usr/local/go/src/go-one/go-app

ADD . /usr/local/go/src/go-one/go-app

RUN go get gopkg.in/ini.v1 && \
  go get github.com/google/uuid && \
  go get github.com/mattn/go-sqlite3

CMD ["go", "run", "main.go"]
