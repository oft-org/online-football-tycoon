# Etapa de build
FROM golang:1.24 AS builder

WORKDIR /app

# Copiar go.mod y go.sum
COPY go.mod ./
COPY go.sum ./
RUN go mod download

# Copiar el código fuente
COPY . .

WORKDIR /app/cmd/api

# Compilar el binario de la aplicación
RUN go build -o /app/app .

# Etapa final
FROM debian:bookworm-slim

# Instalar dependencias necesarias
RUN apt-get update && apt-get install -y \
    bash \
    curl \
    ca-certificates \
    && rm -rf /var/lib/apt/lists/*

# Descargar y preparar la herramienta de migración
RUN curl -L https://github.com/golang-migrate/migrate/releases/download/v4.15.2/migrate.linux-amd64 \
    -o /migrate && chmod +x /migrate

# Copiar el binario construido en la etapa anterior
COPY --from=builder /app/app /app/app

# Copiar scripts necesarios
COPY scripts/wait-for-it.sh /wait-for-it.sh
COPY scripts/migrate.sh /migrate.sh
RUN chmod +x /wait-for-it.sh /migrate.sh

CMD ["/app/app"]
