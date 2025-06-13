# Construcción
FROM golang:1.24.4 as builder

WORKDIR /app
COPY . .

RUN go mod download
RUN go build -o orbe-api ./cmd/orbe.go

# Ejecución
FROM debian:bullseye-slim

WORKDIR /app
COPY --from=builder /app/orbe-api .
COPY migrations ./migrations

EXPOSE 8080

CMD ["./orbe-api"]
