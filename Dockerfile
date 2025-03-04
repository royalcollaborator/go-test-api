FROM golang:1.23.6-alpine

RUN mkdir /app

ADD . /app

WORKDIR /app

RUN go build -o main main.go

CMD ["/app/main"]