# Etapa 1: Build
FROM golang:1.22.5-alpine AS builder

# Configura a variavel de ambiente
ENV GO111MODULE=on

# Cria e define o diretório de trabalho
WORKDIR /app

# Copia o go.mod e o go.sum
COPY go.mod go.sum ./

# Baixa as dependências
RUN go mod download

# Copia o código fonte para o container
COPY . .

# Compila o aplicativo
RUN go build -o main .

# Etapa 2: Run
FROM alpine:latest

# Cria um diretório para a aplicação
WORKDIR /root/

# Copia o binário do build anterior
COPY --from=builder /app/main .

# Define variáveis de ambiente padrão
ENV DB_HOST=192.168.1.88 \
    DB_USER=user \
    DB_PASSWORD=password \
    DB_NAME=dbname \
    DB_PORT=3306

# Expõe a porta em que a aplicação será executada
EXPOSE 8080

# Comando para executar a aplicação
CMD ["./main"]