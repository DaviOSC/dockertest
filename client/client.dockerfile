FROM golang:latest
WORKDIR /app
COPY go.mod .
COPY client.go .
RUN go build -o client .
CMD ["./client"]