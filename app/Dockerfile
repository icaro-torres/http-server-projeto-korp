FROM golang:1.26-alpine3.23 AS builder

WORKDIR /app

COPY go.mod ./
RUN go mod download

COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -o http-server-projeto-korp .

FROM alpine:3.24.1

WORKDIR /root/

COPY --from=builder /app/http-server-projeto-korp .

EXPOSE 8080
CMD ["./http-server-projeto-korp"]