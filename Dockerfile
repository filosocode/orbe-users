# Etapa única: construir y ejecutar en la misma imagen
FROM golang:1.24.4

WORKDIR /app

COPY . .

RUN go mod download
RUN go build -o orbe-api ./cmd/orbe.go

EXPOSE 8080

CMD ["./orbe-api"]
