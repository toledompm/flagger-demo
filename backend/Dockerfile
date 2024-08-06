FROM golang:1.22-alpine AS builder

ARG TAG

WORKDIR /app

COPY go.mod go.mod
COPY go.sum go.sum

RUN go mod download

COPY main.go main.go
COPY domain domain

RUN CGO_ENABLED=0 go build -o /app/main -ldflags "-X main.chosenColor=$TAG" .

FROM alpine:3.12

COPY --from=builder /app/main /app/main

CMD ["/app/main"]
