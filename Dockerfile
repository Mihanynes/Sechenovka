FROM golang:latest

# Установите необходимые библиотеки
RUN apt-get update && apt-get install -y libsqlite3-dev gcc

WORKDIR /app
COPY go.mod ./
COPY . .

# Выполните сборку
RUN go build -o main main.go

ENTRYPOINT ["./main"]