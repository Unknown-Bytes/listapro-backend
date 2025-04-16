# Estágio de build
FROM golang:1.23-alpine AS builder

WORKDIR /app

# Copiar go.mod e go.sum primeiro para aproveitar o cache do Docker
COPY go.mod go.sum ./
RUN go mod download

# Copiar o código fonte
COPY . .

RUN go mod tidy

# Compilar a aplicação
RUN CGO_ENABLED=0 GOOS=linux go build -o backend main.go

# Estágio final
FROM alpine:latest

RUN apk --no-cache add ca-certificates

WORKDIR /root/

# Copiar o binário compilado do estágio de build
COPY --from=builder /app/backend .

# Expor a porta que o servidor irá usar
EXPOSE 8080

# Comando para executar a aplicação
CMD ["./backend"]
