# builder
FROM golang:1.17.5 as builder

WORKDIR /app

COPY . .

RUN go mod download
RUN go build -o api main.go

# runner
FROM debian:11-slim

COPY --from=builder /app/api /app/

WORKDIR /app

EXPOSE 8080

CMD ["./api"]