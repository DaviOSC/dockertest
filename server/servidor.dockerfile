# Usa a imagem oficial do Go
FROM golang:latest

WORKDIR /app

COPY server.go .

RUN go build -o servidor server.go

EXPOSE 8080

CMD ["./servidor"]
