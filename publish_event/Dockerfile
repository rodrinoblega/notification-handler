FROM golang:1.23 AS builder

WORKDIR /app

COPY ../go.mod go.sum ./
RUN go mod download

COPY ./publish_event/main.go /app/
COPY ../config /app/config
COPY ../src /app/src

# Instalar librerías necesarias para confluent-kafka-go
RUN apt-get update && apt-get install -y librdkafka-dev

# Construir el binario (sin deshabilitar CGO)
RUN go build -o app .

# Imagen final
FROM debian:latest

WORKDIR /app

COPY --from=builder /app/app /app/app
COPY --from=builder /app/config /app/config
COPY --from=builder /app/src /app/src

RUN chmod +x /app/app

ENTRYPOINT ["/app/app"]
