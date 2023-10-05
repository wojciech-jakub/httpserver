
FROM golang:latest AS builder

WORKDIR /app

COPY . .

RUN go mod download

COPY . .

RUN go build -o /hello_go_http

EXPOSE 8080

ENTRYPOINT ["/hello_go_http"]