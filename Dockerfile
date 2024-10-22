FROM golang:latest

# Устанавливаем необходимые библиотеки
RUN apt-get update && apt-get install -y gcc libsqlite3-dev

WORKDIR /app
COPY go.mod ./
COPY . .
RUN go build -o main main.go

ENTRYPOINT ["./main"]