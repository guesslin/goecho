FROM golang:latest

MAINTAINER guesslin1986@gmail.com

RUN go get -u github.com/kardianos/govendor

RUN mkdir -p /go/src/app/vendor
WORKDIR /go/src/app
COPY vendor/vendor.json /go/src/app/vendor/
RUN govendor sync

COPY . /go/src/app
RUN go-wrapper install

CMD ["go-wrapper", "run", "-port", "8080", "-ip", "0.0.0.0"]

EXPOSE 8081
