FROM golang:1.24 AS builder
WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -o /app/server ./main.go
FROM alpine:latest
RUN apk add --no-cache bash

WORKDIR /root/
COPY --from=builder /app/server .

EXPOSE 8080

CMD ["./server"]