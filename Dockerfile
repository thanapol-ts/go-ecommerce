FROM golang:1.19-alpine3.16 AS builder

WORKDIR /app

COPY . /app

RUN go build -v -o main .


FROM alpine:3.16
WORKDIR /app
COPY --from=builder /app/main .

EXPOSE 4000

CMD ["/app/main"]