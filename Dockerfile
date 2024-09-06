FROM golang:1.23.1-alpine3.20 AS builder

WORKDIR /
COPY main.go .
COPY go.mod .
COPY go.sum .
RUN go build

FROM alpine:3.20.2
WORKDIR /
COPY --from=builder /ssh-portfolio .
COPY markdown ./markdown

EXPOSE 23234/udp

ENTRYPOINT ["/ssh-portfolio"]

