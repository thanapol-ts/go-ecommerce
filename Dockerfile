FROM golang:latest AS builder

WORKDIR /app

COPY . /app

RUN go build -o main .

CMD ["/app/main"]

